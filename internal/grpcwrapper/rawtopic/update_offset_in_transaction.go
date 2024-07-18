package rawtopic

import (
	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Topic"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/grpcwrapper/rawtopic/rawtopiccommon"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/grpcwrapper/rawydb"
)

type UpdateOffsetsInTransactionRequest struct {
	OperationParams rawydb.OperationParams
	Tx              UpdateOffsetsInTransactionRequest_TransactionIdentity
	Topics          []UpdateOffsetsInTransactionRequest_TopicOffsets
	Consumer        string
}

func (r *UpdateOffsetsInTransactionRequest) ToProto() *Ydb_Topic.UpdateOffsetsInTransactionRequest {
	req := &Ydb_Topic.UpdateOffsetsInTransactionRequest{
		OperationParams: r.OperationParams.ToProto(),
		Tx: &Ydb_Topic.TransactionIdentity{
			Id:      r.Tx.ID,
			Session: r.Tx.Session,
		},
		Consumer: r.Consumer,
	}

	req.Topics = make([]*Ydb_Topic.UpdateOffsetsInTransactionRequest_TopicOffsets, len(r.Topics))
	for topicIndex, topic := range r.Topics {
		offsets := &Ydb_Topic.UpdateOffsetsInTransactionRequest_TopicOffsets{
			Path: topic.Path,
		}

		offsets.Partitions = make([]*Ydb_Topic.UpdateOffsetsInTransactionRequest_TopicOffsets_PartitionOffsets, len(topic.Partitions))
		for partitionIndex, partition := range topic.Partitions {
			protoPartition := &Ydb_Topic.UpdateOffsetsInTransactionRequest_TopicOffsets_PartitionOffsets{
				PartitionId: partition.PartitionID,
			}

			protoPartition.PartitionOffsets = make([]*Ydb_Topic.OffsetsRange, len(partition.PartitionOffsets))

			for offsetIndex, offset := range partition.PartitionOffsets {
				protoPartition.PartitionOffsets[offsetIndex] = offset.ToProto()
			}

			offsets.Partitions[partitionIndex] = protoPartition
		}

		req.Topics[topicIndex] = offsets
	}

	return req
}

type UpdateOffsetsInTransactionRequest_TransactionIdentity struct {
	ID      string
	Session string
}

type UpdateOffsetsInTransactionRequest_TopicOffsets struct {
	Path       string // Topic path
	Partitions []UpdateOffsetsInTransactionRequest_PartitionOffsets
}

type UpdateOffsetsInTransactionRequest_PartitionOffsets struct {
	PartitionID      int64
	PartitionOffsets []rawtopiccommon.OffsetRange
}
