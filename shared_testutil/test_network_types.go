package shared_testutil

import (
	"errors"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/require"

	rm "github.com/LIUYAN-0626/test01-go-fil-markets/retrievalmarket"
	rmnet "github.com/LIUYAN-0626/test01-go-fil-markets/retrievalmarket/network"
	smnet "github.com/LIUYAN-0626/test01-go-fil-markets/storagemarket/network"
)

// QueryReader is a function to mock reading queries.
type QueryReader func() (rm.Query, error)

// QueryResponseReader is a function to mock reading query responses.
type QueryResponseReader func() (rm.QueryResponse, error)

// QueryResponseWriter is a function to mock writing query responses.
type QueryResponseWriter func(rm.QueryResponse) error

// QueryWriter is a function to mock writing queries.
type QueryWriter func(rm.Query) error

// TestRetrievalQueryStream is a retrieval query stream with predefined
// stubbed behavior.
type TestRetrievalQueryStream struct {
	p          peer.ID
	reader     QueryReader
	respReader QueryResponseReader
	respWriter QueryResponseWriter
	writer     QueryWriter
}

// TestQueryStreamParams are parameters used to setup a TestRetrievalQueryStream.
// All parameters except the peer ID are optional.
type TestQueryStreamParams struct {
	PeerID     peer.ID
	Reader     QueryReader
	RespReader QueryResponseReader
	RespWriter QueryResponseWriter
	Writer     QueryWriter
}

// NewTestRetrievalQueryStream returns a new TestRetrievalQueryStream with the
// behavior specified by the paramaters, or default behaviors if not specified.
func NewTestRetrievalQueryStream(params TestQueryStreamParams) rmnet.RetrievalQueryStream {
	stream := TestRetrievalQueryStream{
		p:          params.PeerID,
		reader:     TrivialQueryReader,
		respReader: TrivialQueryResponseReader,
		respWriter: TrivialQueryResponseWriter,
		writer:     TrivialQueryWriter,
	}
	if params.Reader != nil {
		stream.reader = params.Reader
	}
	if params.Writer != nil {
		stream.writer = params.Writer
	}
	if params.RespReader != nil {
		stream.respReader = params.RespReader
	}
	if params.RespWriter != nil {
		stream.respWriter = params.RespWriter
	}
	return &stream
}

// ReadQuery calls the mocked query reader.
func (trqs *TestRetrievalQueryStream) ReadQuery() (rm.Query, error) {
	return trqs.reader()
}

// WriteQuery calls the mocked query writer.
func (trqs *TestRetrievalQueryStream) WriteQuery(newQuery rm.Query) error {
	return trqs.writer(newQuery)
}

// ReadQueryResponse calls the mocked query response reader.
func (trqs *TestRetrievalQueryStream) ReadQueryResponse() (rm.QueryResponse, error) {
	return trqs.respReader()
}

// WriteQueryResponse calls the mocked query response writer.
func (trqs *TestRetrievalQueryStream) WriteQueryResponse(newResp rm.QueryResponse) error {
	return trqs.respWriter(newResp)
}

// Close closes the stream (does nothing for test).
func (trqs *TestRetrievalQueryStream) Close() error { return nil }

// DealProposalReader is a function to mock reading deal proposals.
type DealProposalReader func() (rm.DealProposal, error)

// DealResponseReader is a function to mock reading deal responses.
type DealResponseReader func() (rm.DealResponse, error)

// DealResponseWriter is a function to mock writing deal responses.
type DealResponseWriter func(rm.DealResponse) error

// DealProposalWriter is a function to mock writing deal proposals.
type DealProposalWriter func(rm.DealProposal) error

// DealPaymentReader is a function to mock reading deal payments.
type DealPaymentReader func() (rm.DealPayment, error)

// DealPaymentWriter is a function to mock writing deal payments.
type DealPaymentWriter func(rm.DealPayment) error

// TestRetrievalDealStream is a retrieval deal stream with predefined
// stubbed behavior.
type TestRetrievalDealStream struct {
	p              peer.ID
	proposalReader DealProposalReader
	proposalWriter DealProposalWriter
	responseReader DealResponseReader
	responseWriter DealResponseWriter
	paymentReader  DealPaymentReader
	paymentWriter  DealPaymentWriter
}

// TestDealStreamParams are parameters used to setup a TestRetrievalDealStream.
// All parameters except the peer ID are optional.
type TestDealStreamParams struct {
	PeerID         peer.ID
	ProposalReader DealProposalReader
	ProposalWriter DealProposalWriter
	ResponseReader DealResponseReader
	ResponseWriter DealResponseWriter
	PaymentReader  DealPaymentReader
	PaymentWriter  DealPaymentWriter
}

// NewTestRetrievalDealStream returns a new TestRetrievalDealStream with the
// behavior specified by the paramaters, or default behaviors if not specified.
func NewTestRetrievalDealStream(params TestDealStreamParams) rmnet.RetrievalDealStream {
	stream := TestRetrievalDealStream{
		p:              params.PeerID,
		proposalReader: TrivialDealProposalReader,
		proposalWriter: TrivialDealProposalWriter,
		responseReader: TrivialDealResponseReader,
		responseWriter: TrivialDealResponseWriter,
		paymentReader:  TrivialDealPaymentReader,
		paymentWriter:  TrivialDealPaymentWriter,
	}
	if params.ProposalReader != nil {
		stream.proposalReader = params.ProposalReader
	}
	if params.ProposalWriter != nil {
		stream.proposalWriter = params.ProposalWriter
	}
	if params.ResponseReader != nil {
		stream.responseReader = params.ResponseReader
	}
	if params.ResponseWriter != nil {
		stream.responseWriter = params.ResponseWriter
	}
	if params.PaymentReader != nil {
		stream.paymentReader = params.PaymentReader
	}
	if params.PaymentWriter != nil {
		stream.paymentWriter = params.PaymentWriter
	}
	return &stream
}

// ReadDealProposal calls the mocked deal proposal reader function.
func (trds *TestRetrievalDealStream) ReadDealProposal() (rm.DealProposal, error) {
	return trds.proposalReader()
}

// WriteDealProposal calls the mocked deal proposal writer function.
func (trds *TestRetrievalDealStream) WriteDealProposal(dealProposal rm.DealProposal) error {
	return trds.proposalWriter(dealProposal)
}

// ReadDealResponse calls the mocked deal response reader function.
func (trds *TestRetrievalDealStream) ReadDealResponse() (rm.DealResponse, error) {
	return trds.responseReader()
}

// WriteDealResponse calls the mocked deal response writer function.
func (trds *TestRetrievalDealStream) WriteDealResponse(dealResponse rm.DealResponse) error {
	return trds.responseWriter(dealResponse)
}

// ReadDealPayment calls the mocked deal payment reader function.
func (trds *TestRetrievalDealStream) ReadDealPayment() (rm.DealPayment, error) {
	return trds.paymentReader()
}

// WriteDealPayment calls the mocked deal payment writer function.
func (trds *TestRetrievalDealStream) WriteDealPayment(dealPayment rm.DealPayment) error {
	return trds.paymentWriter(dealPayment)
}

// Receiver returns the other peer
func (trds TestRetrievalDealStream) Receiver() peer.ID { return trds.p }

// Close closes the stream (does nothing for mocked stream)
func (trds TestRetrievalDealStream) Close() error { return nil }

// QueryStreamBuilder is a function that builds retrieval query streams.
type QueryStreamBuilder func(peer.ID) (rmnet.RetrievalQueryStream, error)

// DealStreamBuilder if a function that builds retrieval deal streams
type DealStreamBuilder func(peer.ID) (rmnet.RetrievalDealStream, error)

// TestRetrievalMarketNetwork is a test network that has stubbed behavior
// for testing the retrieval market implementation
type TestRetrievalMarketNetwork struct {
	receiver  rmnet.RetrievalReceiver
	dsbuilder DealStreamBuilder
	qsbuilder QueryStreamBuilder
}

// TestNetworkParams are parameters for setting up a test network. All
// parameters other than the receiver are optional
type TestNetworkParams struct {
	DealStreamBuilder  DealStreamBuilder
	QueryStreamBuilder QueryStreamBuilder
	Receiver           rmnet.RetrievalReceiver
}

// NewTestRetrievalMarketNetwork returns a new TestRetrievalMarketNetwork with the
// behavior specified by the paramaters, or default behaviors if not specified.
func NewTestRetrievalMarketNetwork(params TestNetworkParams) *TestRetrievalMarketNetwork {
	trmn := TestRetrievalMarketNetwork{
		dsbuilder: TrivialNewDealStream,
		qsbuilder: TrivialNewQueryStream,
		receiver:  params.Receiver,
	}
	if params.DealStreamBuilder != nil {
		trmn.dsbuilder = params.DealStreamBuilder
	}
	if params.QueryStreamBuilder != nil {
		trmn.qsbuilder = params.QueryStreamBuilder
	}
	return &trmn
}

// NewQueryStream returns a query stream.
// Note this always returns the same stream.  This is fine for testing for now.
func (trmn *TestRetrievalMarketNetwork) NewQueryStream(id peer.ID) (rmnet.RetrievalQueryStream, error) {
	return trmn.qsbuilder(id)
}

// NewDealStream returns a deal stream
// Note this always returns the same stream.  This is fine for testing for now.
func (trmn *TestRetrievalMarketNetwork) NewDealStream(id peer.ID) (rmnet.RetrievalDealStream, error) {
	return trmn.dsbuilder(id)
}

// SetDelegate sets the market receiver
func (trmn *TestRetrievalMarketNetwork) SetDelegate(r rmnet.RetrievalReceiver) error {
	trmn.receiver = r
	return nil
}

// ReceiveQueryStream simulates receiving a query stream
func (trmn *TestRetrievalMarketNetwork) ReceiveQueryStream(qs rmnet.RetrievalQueryStream) {
	trmn.receiver.HandleQueryStream(qs)
}

// ReceiveDealStream simulates receiving a deal stream
func (trmn *TestRetrievalMarketNetwork) ReceiveDealStream(ds rmnet.RetrievalDealStream) {
	trmn.receiver.HandleDealStream(ds)
}

// StopHandlingRequests sets receiver to nil
func (trmn *TestRetrievalMarketNetwork) StopHandlingRequests() error {
	trmn.receiver = nil
	return nil
}

var _ rmnet.RetrievalMarketNetwork = &TestRetrievalMarketNetwork{}

// Some convenience builders

// FailNewQueryStream always fails
func FailNewQueryStream(peer.ID) (rmnet.RetrievalQueryStream, error) {
	return nil, errors.New("new query stream failed")
}

// FailNewDealStream always fails
func FailNewDealStream(peer.ID) (rmnet.RetrievalDealStream, error) {
	return nil, errors.New("new deal stream failed")
}

// FailQueryReader always fails
func FailQueryReader() (rm.Query, error) {
	return rm.QueryUndefined, errors.New("read query failed")
}

// FailQueryWriter always fails
func FailQueryWriter(rm.Query) error {
	return errors.New("write query failed")
}

// FailResponseReader always fails
func FailResponseReader() (rm.QueryResponse, error) {
	return rm.QueryResponseUndefined, errors.New("query response failed")
}

// FailResponseWriter always fails
func FailResponseWriter(rm.QueryResponse) error {
	return errors.New("write query response failed")
}

// FailDealProposalWriter always fails
func FailDealProposalWriter(rm.DealProposal) error {
	return errors.New("write proposal failed")
}

// FailDealProposalReader always fails
func FailDealProposalReader() (rm.DealProposal, error) {
	return rm.DealProposalUndefined, errors.New("read proposal failed")
}

// FailDealResponseWriter always fails
func FailDealResponseWriter(rm.DealResponse) error {
	return errors.New("write proposal failed")
}

// FailDealResponseReader always fails
func FailDealResponseReader() (rm.DealResponse, error) {
	return rm.DealResponseUndefined, errors.New("write proposal failed")
}

// FailDealPaymentWriter always fails
func FailDealPaymentWriter(rm.DealPayment) error {
	return errors.New("write proposal failed")
}

// FailDealPaymentReader always fails
func FailDealPaymentReader() (rm.DealPayment, error) {
	return rm.DealPaymentUndefined, errors.New("write proposal failed")
}

// TrivialNewQueryStream succeeds trivially, returning an empty query stream.
func TrivialNewQueryStream(p peer.ID) (rmnet.RetrievalQueryStream, error) {
	return NewTestRetrievalQueryStream(TestQueryStreamParams{PeerID: p}), nil
}

// ExpectPeerOnQueryStreamBuilder fails if the peer used does not match the expected peer
func ExpectPeerOnQueryStreamBuilder(t *testing.T, expectedPeer peer.ID, qb QueryStreamBuilder, msgAndArgs ...interface{}) QueryStreamBuilder {
	return func(p peer.ID) (rmnet.RetrievalQueryStream, error) {
		require.Equal(t, expectedPeer, p, msgAndArgs...)
		return qb(p)
	}
}

// TrivialNewDealStream succeeds trivially, returning an empty deal stream.
func TrivialNewDealStream(p peer.ID) (rmnet.RetrievalDealStream, error) {
	return NewTestRetrievalDealStream(TestDealStreamParams{PeerID: p}), nil
}

// TrivialQueryReader succeeds trivially, returning an empty query.
func TrivialQueryReader() (rm.Query, error) {
	return rm.Query{}, nil
}

// TrivialQueryResponseReader succeeds trivially, returning an empty query response.
func TrivialQueryResponseReader() (rm.QueryResponse, error) {
	return rm.QueryResponse{}, nil
}

// TrivialQueryWriter succeeds trivially, returning no error.
func TrivialQueryWriter(rm.Query) error {
	return nil
}

// TrivialQueryResponseWriter succeeds trivially, returning no error.
func TrivialQueryResponseWriter(rm.QueryResponse) error {
	return nil
}

// TrivialDealProposalReader succeeds trivially, returning an empty proposal.
func TrivialDealProposalReader() (rm.DealProposal, error) {
	return rm.DealProposal{}, nil
}

// TrivialDealResponseReader succeeds trivially, returning an empty deal response.
func TrivialDealResponseReader() (rm.DealResponse, error) {
	return rm.DealResponse{}, nil
}

// TrivialDealPaymentReader succeeds trivially, returning an empty deal payment.
func TrivialDealPaymentReader() (rm.DealPayment, error) {
	return rm.DealPayment{}, nil
}

// TrivialDealProposalWriter succeeds trivially, returning no error.
func TrivialDealProposalWriter(rm.DealProposal) error {
	return nil
}

// TrivialDealResponseWriter succeeds trivially, returning no error.
func TrivialDealResponseWriter(rm.DealResponse) error {
	return nil
}

// TrivialDealPaymentWriter succeeds trivially, returning no error.
func TrivialDealPaymentWriter(rm.DealPayment) error {
	return nil
}

// StubbedQueryReader returns the given query when called
func StubbedQueryReader(query rm.Query) QueryReader {
	return func() (rm.Query, error) {
		return query, nil
	}
}

// StubbedQueryResponseReader returns the given query response when called
func StubbedQueryResponseReader(queryResponse rm.QueryResponse) QueryResponseReader {
	return func() (rm.QueryResponse, error) {
		return queryResponse, nil
	}
}

// ExpectQueryWriter will fail if the written query and expected query don't match
func ExpectQueryWriter(t *testing.T, expectedQuery rm.Query, msgAndArgs ...interface{}) QueryWriter {
	return func(query rm.Query) error {
		require.Equal(t, expectedQuery, query, msgAndArgs...)
		return nil
	}
}

// ExpectQueryResponseWriter will fail if the written query response and expected query response don't match
func ExpectQueryResponseWriter(t *testing.T, expectedQueryResponse rm.QueryResponse, msgAndArgs ...interface{}) QueryResponseWriter {
	return func(queryResponse rm.QueryResponse) error {
		require.Equal(t, expectedQueryResponse, queryResponse, msgAndArgs...)
		return nil
	}
}

// ExpectDealResponseWriter will fail if the written query and expected query don't match
func ExpectDealResponseWriter(t *testing.T, expectedDealResponse rm.DealResponse, msgAndArgs ...interface{}) DealResponseWriter {
	return func(dealResponse rm.DealResponse) error {
		require.Equal(t, expectedDealResponse, dealResponse, msgAndArgs...)
		return nil
	}
}

// QueryReadWriter will read only if something is written, otherwise it errors
func QueryReadWriter() (QueryReader, QueryWriter) {
	var q rm.Query
	var written bool
	queryRead := func() (rm.Query, error) {
		if written {
			return q, nil
		}
		return rm.QueryUndefined, errors.New("Unable to read value")
	}
	queryWrite := func(wq rm.Query) error {
		q = wq
		written = true
		return nil
	}
	return queryRead, queryWrite
}

// QueryResponseReadWriter will read only if something is written, otherwise it errors
func QueryResponseReadWriter() (QueryResponseReader, QueryResponseWriter) {
	var q rm.QueryResponse
	var written bool
	queryResponseRead := func() (rm.QueryResponse, error) {
		if written {
			return q, nil
		}
		return rm.QueryResponseUndefined, errors.New("Unable to read value")
	}
	queryResponseWrite := func(wq rm.QueryResponse) error {
		q = wq
		written = true
		return nil
	}
	return queryResponseRead, queryResponseWrite
}

// StubbedDealProposalReader returns the given proposal when called
func StubbedDealProposalReader(proposal rm.DealProposal) DealProposalReader {
	return func() (rm.DealProposal, error) {
		return proposal, nil
	}
}

// StubbedDealResponseReader returns the given deal response when called
func StubbedDealResponseReader(response rm.DealResponse) DealResponseReader {
	return func() (rm.DealResponse, error) {
		return response, nil
	}
}

// StubbedDealPaymentReader returns the given deal payment when called
func StubbedDealPaymentReader(payment rm.DealPayment) DealPaymentReader {
	return func() (rm.DealPayment, error) {
		return payment, nil
	}
}

// StorageDealProposalReader is a function to mock reading deal proposals.
type StorageDealProposalReader func() (smnet.Proposal, error)

// StorageDealResponseReader is a function to mock reading deal responses.
type StorageDealResponseReader func() (smnet.SignedResponse, error)

// StorageDealResponseWriter is a function to mock writing deal responses.
type StorageDealResponseWriter func(smnet.SignedResponse) error

// StorageDealProposalWriter is a function to mock writing deal proposals.
type StorageDealProposalWriter func(smnet.Proposal) error

// TestStorageDealStream is a retrieval deal stream with predefined
// stubbed behavior.
type TestStorageDealStream struct {
	p              peer.ID
	proposalReader StorageDealProposalReader
	proposalWriter StorageDealProposalWriter
	responseReader StorageDealResponseReader
	responseWriter StorageDealResponseWriter
	tags           map[string]struct{}
}

// TestStorageDealStreamParams are parameters used to setup a TestStorageDealStream.
// All parameters except the peer ID are optional.
type TestStorageDealStreamParams struct {
	PeerID         peer.ID
	ProposalReader StorageDealProposalReader
	ProposalWriter StorageDealProposalWriter
	ResponseReader StorageDealResponseReader
	ResponseWriter StorageDealResponseWriter
}

// NewTestStorageDealStream returns a new TestStorageDealStream with the
// behavior specified by the paramaters, or default behaviors if not specified.
func NewTestStorageDealStream(params TestStorageDealStreamParams) *TestStorageDealStream {
	stream := TestStorageDealStream{
		p:              params.PeerID,
		proposalReader: TrivialStorageDealProposalReader,
		proposalWriter: TrivialStorageDealProposalWriter,
		responseReader: TrivialStorageDealResponseReader,
		responseWriter: TrivialStorageDealResponseWriter,
		tags:           make(map[string]struct{}),
	}
	if params.ProposalReader != nil {
		stream.proposalReader = params.ProposalReader
	}
	if params.ProposalWriter != nil {
		stream.proposalWriter = params.ProposalWriter
	}
	if params.ResponseReader != nil {
		stream.responseReader = params.ResponseReader
	}
	if params.ResponseWriter != nil {
		stream.responseWriter = params.ResponseWriter
	}
	return &stream
}

// ReadDealProposal calls the mocked deal proposal reader function.
func (tsds *TestStorageDealStream) ReadDealProposal() (smnet.Proposal, error) {
	return tsds.proposalReader()
}

// WriteDealProposal calls the mocked deal proposal writer function.
func (tsds *TestStorageDealStream) WriteDealProposal(dealProposal smnet.Proposal) error {
	return tsds.proposalWriter(dealProposal)
}

// ReadDealResponse calls the mocked deal response reader function.
func (tsds *TestStorageDealStream) ReadDealResponse() (smnet.SignedResponse, error) {
	return tsds.responseReader()
}

// WriteDealResponse calls the mocked deal response writer function.
func (tsds *TestStorageDealStream) WriteDealResponse(dealResponse smnet.SignedResponse) error {
	return tsds.responseWriter(dealResponse)
}

// RemotePeer returns the other peer
func (tsds TestStorageDealStream) RemotePeer() peer.ID { return tsds.p }

// Close closes the stream (does nothing for mocked stream)
func (tsds TestStorageDealStream) Close() error { return nil }

// TagProtectedConnection preserves this connection as higher priority than others
func (tsds TestStorageDealStream) TagProtectedConnection(identifier string) {
	tsds.tags[identifier] = struct{}{}
}

// UntagProtectedConnection removes the given tag on this connection, increasing
// the likelyhood it will be cleaned up
func (tsds TestStorageDealStream) UntagProtectedConnection(identifier string) {
	delete(tsds.tags, identifier)
}

// AssertConnectionTagged verifies a connection was tagged with the given identifier
func (tsds TestStorageDealStream) AssertConnectionTagged(t *testing.T, identifier string) {
	_, ok := tsds.tags[identifier]
	require.True(t, ok)
}

// TrivialStorageDealProposalReader succeeds trivially, returning an empty proposal.
func TrivialStorageDealProposalReader() (smnet.Proposal, error) {
	return smnet.Proposal{}, nil
}

// TrivialStorageDealResponseReader succeeds trivially, returning an empty deal response.
func TrivialStorageDealResponseReader() (smnet.SignedResponse, error) {
	return smnet.SignedResponse{}, nil
}

// TrivialStorageDealProposalWriter succeeds trivially, returning no error.
func TrivialStorageDealProposalWriter(smnet.Proposal) error {
	return nil
}

// TrivialStorageDealResponseWriter succeeds trivially, returning no error.
func TrivialStorageDealResponseWriter(smnet.SignedResponse) error {
	return nil
}

// StubbedStorageProposalReader returns the given proposal when called
func StubbedStorageProposalReader(proposal smnet.Proposal) StorageDealProposalReader {
	return func() (smnet.Proposal, error) {
		return proposal, nil
	}
}

// StubbedStorageResponseReader returns the given deal response when called
func StubbedStorageResponseReader(response smnet.SignedResponse) StorageDealResponseReader {
	return func() (smnet.SignedResponse, error) {
		return response, nil
	}
}

// FailStorageProposalWriter always fails
func FailStorageProposalWriter(smnet.Proposal) error {
	return errors.New("write proposal failed")
}

// FailStorageProposalReader always fails
func FailStorageProposalReader() (smnet.Proposal, error) {
	return smnet.ProposalUndefined, errors.New("read proposal failed")
}

// FailStorageResponseWriter always fails
func FailStorageResponseWriter(smnet.SignedResponse) error {
	return errors.New("write proposal failed")
}

// FailStorageResponseReader always fails
func FailStorageResponseReader() (smnet.SignedResponse, error) {
	return smnet.SignedResponseUndefined, errors.New("read response failed")
}

// TestPeerResolver provides a fake retrievalmarket PeerResolver
type TestPeerResolver struct {
	Peers         []rm.RetrievalPeer
	ResolverError error
}

func (tpr TestPeerResolver) GetPeers(cid.Cid) ([]rm.RetrievalPeer, error) {
	return tpr.Peers, tpr.ResolverError
}

var _ rm.PeerResolver = &TestPeerResolver{}
