//go:generate mockgen  -destination=./mocks/stream.go -package=mocks github.com/blocknative/dreamboat/pkg/stream Pubsub,Datastore

package stream

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lthibault/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/blocknative/dreamboat/structs"
	"github.com/blocknative/dreamboat/structs/forks/bellatrix"
	"github.com/blocknative/dreamboat/structs/forks/capella"
)

var (
	CacheTopic         = "/block/cache"
	BidTopic           = "/block/bid"
	SlotDeliveredTopic = "/slot/delivered"
)

type Pubsub interface {
	Publish(context.Context, string, []byte) error
	Subscribe(context.Context, string) chan []byte
}

type StreamConfig struct {
	Logger          log.Logger
	ID              string
	PubsubTopic     string // pubsub topic name for block submissions
	StreamQueueSize int
}

type State interface {
	ForkVersion(epoch structs.Slot) structs.ForkVersion
	HeadSlot() structs.Slot
}

type Client struct {
	Pubsub Pubsub

	builderBidIn     chan []byte
	builderBidOut    chan structs.BuilderBidExtended
	cacheIn          chan []byte
	cacheOut         chan structs.BlockAndTraceExtended
	slotDeliveredIn  chan []byte
	slotDeliveredOut chan uint64

	Config StreamConfig
	Logger log.Logger

	m StreamMetrics

	//slotDelivered chan structs.Slot

	st State
}

func NewClient(ps Pubsub, st State, cfg StreamConfig) *Client {
	s := Client{
		Pubsub: ps,
		st:     st,

		builderBidIn:     make(chan []byte, cfg.StreamQueueSize),
		builderBidOut:    make(chan structs.BuilderBidExtended, cfg.StreamQueueSize),
		cacheIn:          make(chan []byte, cfg.StreamQueueSize),
		cacheOut:         make(chan structs.BlockAndTraceExtended, cfg.StreamQueueSize),
		slotDeliveredIn:  make(chan []byte, cfg.StreamQueueSize),
		slotDeliveredOut: make(chan uint64, cfg.StreamQueueSize),

		Config: cfg,
		Logger: cfg.Logger.WithField("relay-service", "stream").WithField("type", "redis"),
	}

	s.initMetrics()

	return &s
}

func (s *Client) RunSubscriberParallel(ctx context.Context, num uint) error {
	s.builderBidIn = s.Pubsub.Subscribe(ctx, s.Config.PubsubTopic+BidTopic)
	s.cacheIn = s.Pubsub.Subscribe(ctx, s.Config.PubsubTopic+CacheTopic)
	s.slotDeliveredIn = s.Pubsub.Subscribe(ctx, s.Config.PubsubTopic+SlotDeliveredTopic)

	for i := uint(0); i < num; i++ {
		go s.RunCacheSubscriber(ctx)
		go s.RunBuilderBidSubscriber(ctx)
	}

	go s.RunSlotDeliveredSubscriber(ctx)

	return nil
}

func (s *Client) BlockCache() <-chan structs.BlockAndTraceExtended {
	return s.cacheOut
}

func (s *Client) RunCacheSubscriber(ctx context.Context) error {
	l := s.Logger.WithField("method", "runCacheSubscriber")
	var bbt structs.BlockAndTraceExtended

	for raw := range s.cacheIn {
		sData, err := s.decode(raw)
		if err != nil {
			l.WithError(err).Warn("failed to decode cache wrapper")
			continue
		}

		if sData.Meta().Source == s.Config.ID {
			continue
		}

		switch forkEncoding := sData.Meta().ForkEncoding; forkEncoding {
		case BellatrixJson:
			var bbbt bellatrix.BlockBidAndTrace
			if err := json.Unmarshal(sData.Data(), &bbbt); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode cache")
				continue
			}
			bbt = &bbbt
		case CapellaJson:
			var cbbt capella.BlockAndTraceExtended
			if err := json.Unmarshal(sData.Data(), &cbbt); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode cache")
				continue
			}
			bbt = &cbbt
		case CapellaSSZ:
			var cbbt capella.BlockAndTraceExtended
			if err := cbbt.UnmarshalSSZ(sData.Data()); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode cache")
				continue
			}
			bbt = &cbbt
		default:
			l.WithField("forkEncoding", forkEncoding).Warn("unkown cache forkEncoding")
			continue
		}

		s.m.StreamRecvCounter.WithLabelValues("cache").Inc()
		select {
		case s.cacheOut <- bbt:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return ctx.Err()
}

func (s *Client) BuilderBid() <-chan structs.BuilderBidExtended {
	return s.builderBidOut
}

func (s *Client) RunBuilderBidSubscriber(ctx context.Context) error {
	l := s.Logger.WithField("method", "runBuilderBidSubscriber")
	var bb structs.BuilderBidExtended

	for raw := range s.builderBidIn {
		sData, err := s.decode(raw)
		if err != nil {
			l.WithError(err).Warn("failed to decode builder bid  wrapper")
			continue
		}

		if sData.Meta().Source == s.Config.ID {
			continue
		}

		switch forkEncoding := sData.Meta().ForkEncoding; forkEncoding {
		case BellatrixJson:
			var bbb bellatrix.BuilderBidExtended
			if err := json.Unmarshal(sData.Data(), &bbb); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode builder bid")
				continue
			}
			bb = &bbb
		case CapellaJson:
			var cbb capella.BuilderBidExtended
			if err := json.Unmarshal(sData.Data(), &cbb); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode builder bid")
				continue
			}
			bb = &cbb
		case CapellaSSZ:
			var cbb capella.BuilderBidExtended
			if err := cbb.UnmarshalSSZ(sData.Data()); err != nil {
				l.WithError(err).WithField("forkEncoding", forkEncoding).Warn("failed to decode builder bid")
				continue
			}
			bb = &cbb
		default:
			l.WithField("forkEncoding", forkEncoding).Warn("unkown builder bid forkEncoding")
			continue
		}

		s.m.StreamRecvCounter.WithLabelValues("bid").Inc()
		select {
		case s.builderBidOut <- bb:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return ctx.Err()
}

func (s *Client) RunSlotDeliveredSubscriber(ctx context.Context) error {
	return nil // TODO
}

func (s *Client) PublishBuilderBid(ctx context.Context, bid structs.BuilderBidExtended) error {
	timer0 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishBuilderBid", "all"))

	timer1 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishBuilderBid", "encode"))
	forkEncoding := toBidFormat(s.st.ForkVersion(structs.Slot(bid.Slot())))
	b, err := s.encode(bid, forkEncoding)
	if err != nil {
		timer1.ObserveDuration()
		return fmt.Errorf("fail to encode encode and stream block: %w", err)
	}
	timer1.ObserveDuration()

	timer2 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishBuilderBid", "publish"))
	if err := s.Pubsub.Publish(ctx, s.Config.PubsubTopic+BidTopic, b); err != nil {
		return fmt.Errorf("fail to encode encode and stream block: %w", err)
	}
	timer2.ObserveDuration()

	timer0.ObserveDuration()
	return nil
}

func (s *Client) PublishBlockCache(ctx context.Context, block structs.BlockAndTraceExtended) error {
	timer0 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishCacheBlock", "all"))

	timer1 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishCacheBlock", "encode"))
	forkEncoding := toBlockCacheFormat(s.st.ForkVersion(structs.Slot(block.Slot())))
	b, err := s.encode(block, forkEncoding)
	if err != nil {
		timer1.ObserveDuration()
		return fmt.Errorf("fail to encode cache block: %w", err)
	}
	timer1.ObserveDuration()

	timer2 := prometheus.NewTimer(s.m.Timing.WithLabelValues("publishCacheBlock", "publish"))
	if err := s.Pubsub.Publish(ctx, s.Config.PubsubTopic+CacheTopic, b); err != nil {
		return fmt.Errorf("fail to publish cache block: %w", err)
	}
	timer2.ObserveDuration()

	timer0.ObserveDuration()
	return nil
}

func (s *Client) PublishSlotDelivered(ctx context.Context, slot structs.Slot) error {
	return nil // TODO
}

func (s *Client) encode(data any, fvf ForkVersionFormat) ([]byte, error) {
	var (
		rawData []byte
		err     error
	)

	if fvf == CapellaSSZ {
		enc, ok := data.(EncoderSSZ)
		if !ok {
			return nil, errors.New("capella ssz unable to cast to SSZ encoder")
		}
		rawData, err = enc.MarshalSSZ()
		if err != nil {
			return nil, err
		}
	} else {
		rawData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	item := JsonItem{
		StreamData: rawData,
		StreamMeta: Metadata{Source: s.Config.ID, ForkEncoding: fvf},
	}

	rawItem, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	// encode the varint with a variable size
	varintBytes := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(varintBytes, uint64(CapellaJson))
	varintBytes = varintBytes[:n]

	// append the varint
	return append(varintBytes, rawItem...), nil
}

type StreamData interface {
	Data() []byte
	Meta() Metadata
}

func (s *Client) decode(b []byte) (StreamData, error) {
	varint, n := binary.Uvarint(b)
	if n <= 0 {
		return nil, ErrDecodeVarint
	}

	b = b[n:]
	forkEncoding := ForkVersionFormat(varint)

	switch forkEncoding {
	case BellatrixJson:
		fallthrough
	case CapellaJson:
		var jsonReq JsonItem
		if err := json.Unmarshal(b, &jsonReq); err != nil {
			return nil, fmt.Errorf("failed to unmarshal json stream data: %w", err)
		}
		return &jsonReq, nil
	}
	return nil, fmt.Errorf("invalid fork version format: %d", forkEncoding)
}

type ForkVersionFormat uint64

const (
	Unknown ForkVersionFormat = iota
	AltairJson
	BellatrixJson
	CapellaJson
	CapellaSSZ
)

var (
	ErrDecodeVarint = errors.New("error decoding varint value")
)

func toBidFormat(fork structs.ForkVersion) ForkVersionFormat {
	switch fork {
	case structs.ForkAltair:
		return AltairJson
	case structs.ForkBellatrix:
		return BellatrixJson
	case structs.ForkCapella:
		return CapellaSSZ
	}
	return Unknown
}

func toBlockCacheFormat(fork structs.ForkVersion) ForkVersionFormat {
	switch fork {
	case structs.ForkAltair:
		return AltairJson
	case structs.ForkBellatrix:
		return BellatrixJson
	case structs.ForkCapella:
		return CapellaSSZ
	}
	return Unknown
}

type EncoderSSZ interface {
	MarshalSSZ() ([]byte, error)
}
