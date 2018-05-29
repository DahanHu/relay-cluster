package shell

import (
	"encoding/binary"
	"encoding/json"

<<<<<<< HEAD
	"github.com/libp2p/go-libp2p-peer"
	pb "github.com/libp2p/go-libp2p-pubsub/pb"
=======
	"github.com/libp2p/go-floodsub"
	"github.com/libp2p/go-libp2p-peer"
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
)

// PubSubRecord is a record received via PubSub.
type PubSubRecord interface {
	// From returns the peer ID of the node that published this record
	From() peer.ID

	// Data returns the data field
	Data() []byte

	// SeqNo is the sequence number of this record
	SeqNo() int64

	//TopicIDs is the list of topics this record belongs to
	TopicIDs() []string
}

<<<<<<< HEAD
type message struct {
	*pb.Message
}

func (m *message) GetFrom() peer.ID {
	return peer.ID(m.Message.GetFrom())
}

type floodsubRecord struct {
	msg *message
=======
type floodsubRecord struct {
	msg *floodsub.Message
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
}

func (r floodsubRecord) From() peer.ID {
	return r.msg.GetFrom()
}

func (r floodsubRecord) Data() []byte {
	return r.msg.GetData()
}

func (r floodsubRecord) SeqNo() int64 {
	return int64(binary.BigEndian.Uint64(r.msg.GetSeqno()))
}

func (r floodsubRecord) TopicIDs() []string {
	return r.msg.GetTopicIDs()
}

///

// PubSubSubscription allow you to receive pubsub records that where published on the network.
type PubSubSubscription struct {
	resp *Response
}

func newPubSubSubscription(resp *Response) *PubSubSubscription {
	sub := &PubSubSubscription{
		resp: resp,
	}

<<<<<<< HEAD
=======
	sub.Next() // skip empty element used for flushing
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
	return sub
}

// Next waits for the next record and returns that.
func (s *PubSubSubscription) Next() (PubSubRecord, error) {
	if s.resp.Error != nil {
		return nil, s.resp.Error
	}

	d := json.NewDecoder(s.resp.Output)

<<<<<<< HEAD
	r := &message{}
=======
	r := &floodsub.Message{}
>>>>>>> 258d5c409a01370dfe542ceadc3d1669659150fe
	err := d.Decode(r)

	return floodsubRecord{msg: r}, err
}

// Cancel cancels the given subscription.
func (s *PubSubSubscription) Cancel() error {
	if s.resp.Output == nil {
		return nil
	}

	return s.resp.Output.Close()
}
