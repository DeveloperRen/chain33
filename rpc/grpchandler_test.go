package rpc

import (
	"testing"

	"gitlab.33.cn/chain33/chain33/types"
	pb "gitlab.33.cn/chain33/chain33/types"
	"golang.org/x/net/context"

	"encoding/hex"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.33.cn/chain33/chain33/client/mocks"
	"google.golang.org/grpc/peer"
)

var (
	g    Grpc
	qapi *mocks.QueueProtocolAPI
)

// Addr is an autogenerated mock type for the Addr type
type Addr struct {
	mock.Mock
}

// Network provides a mock function with given fields:
func (_m *Addr) Network() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Addr) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

func init() {
	//addr := "192.168.1.1"
	//remoteIpWhitelist[addr] = true
	//grpcFuncWhitelist["*"] = true

	qapi = new(mocks.QueueProtocolAPI)
	g.cli.QueueProtocolAPI = qapi
}

func getOkCtx() context.Context {
	addr := new(Addr)
	addr.On("String").Return("192.168.1.1")

	ctx := context.Background()
	pr := &peer.Peer{
		Addr:     addr,
		AuthInfo: nil,
	}
	ctx = peer.NewContext(ctx, pr)
	return ctx
}

//func getNokCtx() context.Context {
//	addr := new(Addr)
//	addr.On("String").Return("192.168.1.0")
//
//	ctx := context.Background()
//	pr := &peer.Peer{
//		Addr:     addr,
//		AuthInfo: nil,
//	}
//	ctx = peer.NewContext(ctx, pr)
//	return ctx
//}

func testSendTransactionOk(t *testing.T) {

	var in *types.Transaction
	reply := &types.Reply{IsOk: true, Msg: nil}
	qapi.On("SendTx", in).Return(reply, nil)
	// NotifySendTxResult called in SendTx
	qapi.On("NotifySendTxResult", &types.ReqNotifySendTxResult{Isok: true, Tx: in}).Return(reply, nil)

	reply, err := g.SendTransaction(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, true, reply.IsOk, "reply should be ok")
}

//func testSendTransactionReject(t *testing.T) {
//	var in *types.Transaction
//
//	_, err := g.SendTransaction(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func TestSendTransaction(t *testing.T) {
	testSendTransactionOk(t)
	//testSendTransactionReject(t)

}

//type fuc func(ctx context.Context, in)

//func testReject(t *testing.T, f func( context.Context, interface{}) (*pb.Reply, error))  {
//func testReject(t *testing.T, f func( context.Context, interface{}) (*pb.Reply, error))  {
//
//	_,err := f(getNokCtx(), nil)
//	assert.EqualError(t,err,"reject","the error should be reject")
//
//
//}

//func testVersionReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.Version(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testVersionOK(t *testing.T) {
	reply := &types.Reply{IsOk: true, Msg: nil}
	qapi.On("Version").Return(reply, nil)
	data, err := g.Version(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, true, data.IsOk, "reply should be ok")

}

func TestVersion(t *testing.T) {
	//testVersionReject(t)
	testVersionOK(t)
}

//func (g *Grpc) GetMemPool(ctx context.Context, in *pb.ReqNil) (*pb.ReplyTxList, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetMempool()
//}

//func testGetMemPoolReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetMemPool(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetMemPoolOK(t *testing.T) {
	qapi.On("GetMempool").Return(nil, nil)
	data, err := g.GetMemPool(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func Test_GetMemPool(t *testing.T) {
	//testGetMemPoolReject(t)
	testGetMemPoolOK(t)
}

//func (g *Grpc) GetLastMemPool(ctx context.Context, in *pb.ReqNil) (*pb.ReplyTxList, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetLastMempool()
//}

//func testGetLastMemPoolReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetLastMemPool(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetLastMemPoolOK(t *testing.T) {
	qapi.On("GetLastMempool").Return(nil, nil)
	data, err := g.GetLastMemPool(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetLastMemPool(t *testing.T) {
	//testGetLastMemPoolReject(t)
	testGetLastMemPoolOK(t)
}

//func (g *Grpc) QueryChain(ctx context.Context, in *pb.Query) (*pb.Reply, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	msg, err := g.cli.Query(in)
//	if err != nil {
//		return nil, err
//	}
//	var reply pb.Reply
//	reply.IsOk = true
//	reply.Msg = pb.Encode(*msg)
//	return &reply, nil
//}

//func testQueryChainReject(t *testing.T) {
//	_, err := g.QueryChain(getNokCtx(), nil)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testQueryChainError(t *testing.T) {
	var in *pb.Query

	qapi.On("Query", in).Return(nil, fmt.Errorf("error")).Once()
	_, err := g.QueryChain(getOkCtx(), in)
	assert.EqualError(t, err, "error", "return error")
}

func testQueryChainOK(t *testing.T) {
	var in *pb.Query
	var msg types.Message
	var req types.ReqString
	req.Data = "msg"
	msg = &req
	qapi.On("Query", in).Return(&msg, nil).Once()
	data, err := g.QueryChain(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, true, data.IsOk, "reply should be ok")
	var decodemsg types.ReqString
	pb.Decode(data.Msg, &decodemsg)
	assert.Equal(t, msg, &decodemsg)

}

func TestQueryChain(t *testing.T) {
	//testQueryChainReject(t)
	testQueryChainError(t)
	testQueryChainOK(t)
}

//func testGetPeerInfoReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetPeerInfo(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetPeerInfoOK(t *testing.T) {
	qapi.On("PeerInfo").Return(nil, nil)
	data, err := g.GetPeerInfo(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetPeerInfo(t *testing.T) {
	//testGetPeerInfoReject(t)
	testGetPeerInfoOK(t)
}

//func (g *Grpc) NetInfo(ctx context.Context, in *pb.ReqNil) (*pb.NodeNetInfo, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetNetInfo()
//}

//func testNetInfoReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.NetInfo(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testNetInfoOK(t *testing.T) {
	qapi.On("GetNetInfo").Return(nil, nil)
	data, err := g.NetInfo(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestNetInfo(t *testing.T) {
	//testNetInfoReject(t)
	testNetInfoOK(t)
}

//func (g *Grpc) GetTicketCount(ctx context.Context, in *pb.ReqNil) (*pb.Int64, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetTicketCount()
//}

//func testGetTicketCountReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetTicketCount(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetTicketCountOK(t *testing.T) {
	qapi.On("GetTicketCount").Return(nil, nil)
	data, err := g.GetTicketCount(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetTicketCount(t *testing.T) {
	//testGetTicketCountReject(t)
	testGetTicketCountOK(t)
}

//func testGetAccountsReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetAccounts(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetAccountsOK(t *testing.T) {
	qapi.On("WalletGetAccountList").Return(nil, nil)
	data, err := g.GetAccounts(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetAccounts(t *testing.T) {
	//testGetAccountsReject(t)
	testGetAccountsOK(t)
}

//func testNewAccountReject(t *testing.T) {
//	var in *pb.ReqNewAccount
//
//	_, err := g.NewAccount(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testNewAccountOK(t *testing.T) {
	var in *pb.ReqNewAccount
	qapi.On("NewAccount", in).Return(nil, nil)
	data, err := g.NewAccount(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestNewAccount(t *testing.T) {
	//testNewAccountReject(t)
	testNewAccountOK(t)
}

//func testWalletTransactionListReject(t *testing.T) {
//	var in *pb.ReqWalletTransactionList
//
//	_, err := g.WalletTransactionList(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testWalletTransactionListOK(t *testing.T) {
	var in *pb.ReqWalletTransactionList
	qapi.On("WalletTransactionList", in).Return(nil, nil)
	data, err := g.WalletTransactionList(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestWalletTransactionList(t *testing.T) {
	//testWalletTransactionListReject(t)
	testWalletTransactionListOK(t)
}

//func testImportPrivKeyReject(t *testing.T) {
//	var in *pb.ReqWalletImportPrivKey
//
//	_, err := g.ImportPrivKey(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testImportPrivKeyOK(t *testing.T) {
	var in *pb.ReqWalletImportPrivKey
	qapi.On("WalletImportprivkey", in).Return(nil, nil)
	data, err := g.ImportPrivKey(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestImportPrivKey(t *testing.T) {
	//testImportPrivKeyReject(t)
	testImportPrivKeyOK(t)
}

//func testSendToAddressReject(t *testing.T) {
//	var in *pb.ReqWalletSendToAddress
//
//	_, err := g.SendToAddress(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSendToAddressOK(t *testing.T) {
	var in *pb.ReqWalletSendToAddress
	qapi.On("WalletSendToAddress", in).Return(nil, nil)
	data, err := g.SendToAddress(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSendToAddress(t *testing.T) {
	//testSendToAddressReject(t)
	testSendToAddressOK(t)
}

//func testSetTxFeeReject(t *testing.T) {
//	var in *pb.ReqWalletSetFee
//
//	_, err := g.SetTxFee(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSetTxFeeOK(t *testing.T) {
	var in *pb.ReqWalletSetFee
	qapi.On("WalletSetFee", in).Return(nil, nil)
	data, err := g.SetTxFee(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSetTxFee(t *testing.T) {
	//testSetTxFeeReject(t)
	testSetTxFeeOK(t)
}

//func testSetLablReject(t *testing.T) {
//	var in *pb.ReqWalletSetLabel
//
//	_, err := g.SetLabl(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSetLablOK(t *testing.T) {
	var in *pb.ReqWalletSetLabel
	qapi.On("WalletSetLabel", in).Return(nil, nil)
	data, err := g.SetLabl(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSetLabl(t *testing.T) {
	//testSetLablReject(t)
	testSetLablOK(t)
}

//func testMergeBalanceReject(t *testing.T) {
//	var in *pb.ReqWalletMergeBalance
//
//	_, err := g.MergeBalance(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testMergeBalanceOK(t *testing.T) {
	var in *pb.ReqWalletMergeBalance
	qapi.On("WalletMergeBalance", in).Return(nil, nil)
	data, err := g.MergeBalance(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestMergeBalance(t *testing.T) {
	//testMergeBalanceReject(t)
	testMergeBalanceOK(t)
}

//func testSetPasswdReject(t *testing.T) {
//	var in *pb.ReqWalletSetPasswd
//
//	_, err := g.SetPasswd(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSetPasswdOK(t *testing.T) {
	var in *pb.ReqWalletSetPasswd
	qapi.On("WalletSetPasswd", in).Return(nil, nil)
	data, err := g.SetPasswd(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSetPasswd(t *testing.T) {
	//testSetPasswdReject(t)
	testSetPasswdOK(t)
}

//func testLockReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.Lock(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testLockOK(t *testing.T) {

	qapi.On("WalletLock").Return(nil, nil)
	data, err := g.Lock(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestLock(t *testing.T) {
	//testLockReject(t)
	testLockOK(t)
}

//func testUnLockReject(t *testing.T) {
//	var in *pb.WalletUnLock
//
//	_, err := g.UnLock(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testUnLockOK(t *testing.T) {
	var in *pb.WalletUnLock
	qapi.On("WalletUnLock", in).Return(nil, nil)
	data, err := g.UnLock(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestUnLock(t *testing.T) {
	//testUnLockReject(t)
	testUnLockOK(t)
}

//func testGenSeedReject(t *testing.T) {
//	var in *pb.GenSeedLang
//
//	_, err := g.GenSeed(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGenSeedOK(t *testing.T) {
	var in *pb.GenSeedLang
	qapi.On("GenSeed", in).Return(nil, nil)
	data, err := g.GenSeed(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGenSeed(t *testing.T) {
	//testGenSeedReject(t)
	testGenSeedOK(t)
}

//func testGetSeedReject(t *testing.T) {
//	var in *pb.GetSeedByPw
//
//	_, err := g.GetSeed(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetSeedOK(t *testing.T) {
	var in *pb.GetSeedByPw
	qapi.On("GetSeed", in).Return(nil, nil)
	data, err := g.GetSeed(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetSeed(t *testing.T) {
	//testGetSeedReject(t)
	testGetSeedOK(t)
}

//func testSaveSeedReject(t *testing.T) {
//	var in *pb.SaveSeedByPw
//
//	_, err := g.SaveSeed(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSaveSeedOK(t *testing.T) {
	var in *pb.SaveSeedByPw
	qapi.On("SaveSeed", in).Return(nil, nil)
	data, err := g.SaveSeed(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSaveSeed(t *testing.T) {
	//testSaveSeedReject(t)
	testSaveSeedOK(t)
}

//func testGetWalletStatusReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetWalletStatus(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetWalletStatusOK(t *testing.T) {
	qapi.On("GetWalletStatus").Return(nil, nil)
	data, err := g.GetWalletStatus(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetWalletStatus(t *testing.T) {
	//testGetWalletStatusReject(t)
	testGetWalletStatusOK(t)
}

//func testSetAutoMiningReject(t *testing.T) {
//	var in *pb.MinerFlag
//
//	_, err := g.SetAutoMining(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testSetAutoMiningOK(t *testing.T) {
	var in *pb.MinerFlag
	qapi.On("WalletAutoMiner", in).Return(nil, nil)
	data, err := g.SetAutoMining(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestSetAutoMining(t *testing.T) {
	//testSetAutoMiningReject(t)
	testSetAutoMiningOK(t)
}

//func testDumpPrivkeyReject(t *testing.T) {
//	var in *pb.ReqStr
//
//	_, err := g.DumpPrivkey(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testDumpPrivkeyOK(t *testing.T) {
	var in *pb.ReqStr
	qapi.On("DumpPrivkey", in).Return(nil, nil)
	data, err := g.DumpPrivkey(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestDumpPrivkey(t *testing.T) {
	//testDumpPrivkeyReject(t)
	testDumpPrivkeyOK(t)
}

//func testCloseTicketsReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.CloseTickets(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testCloseTicketsOK(t *testing.T) {

	qapi.On("CloseTickets").Return(nil, nil)
	data, err := g.CloseTickets(getOkCtx(), nil)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestCloseTickets(t *testing.T) {
	//testCloseTicketsReject(t)
	testCloseTicketsOK(t)
}

//func testGetBlocksReject(t *testing.T) {
//	var in *pb.ReqBlocks
//
//	_, err := g.GetBlocks(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetBlocksError(t *testing.T) {
	var in = pb.ReqBlocks{0, 0, true, []string{""}}

	qapi.On("GetBlocks", &in).Return(nil, fmt.Errorf("error")).Once()
	_, err := g.GetBlocks(getOkCtx(), &in)
	assert.EqualError(t, err, "error", "the error should be error")

}

func testGetBlocksOK(t *testing.T) {
	var in = pb.ReqBlocks{0, 0, true, []string{""}}
	var details types.BlockDetails

	var block = &types.Block{Version: 1}
	var detail = &types.BlockDetail{Block: block, Receipts: nil}
	details.Items = append(details.Items, detail)

	qapi.On("GetBlocks", &in).Return(&details, nil).Once()
	data, err := g.GetBlocks(getOkCtx(), &in)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, true, data.IsOk)

	var details2 types.BlockDetails
	pb.Decode(data.Msg, &details2)
	assert.Equal(t, details, details2)

}

func TestGetBlocks(t *testing.T) {
	//testGetBlocksReject(t)
	testGetBlocksError(t)
	testGetBlocksOK(t)
}

//func (g *Grpc) GetHexTxByHash(ctx context.Context, in *pb.ReqHash) (*pb.HexTx, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	reply, err := g.cli.QueryTx(in)
//	if err != nil {
//		return nil, err
//	}
//	tx := reply.GetTx()
//	if tx == nil {
//		return &pb.HexTx{}, nil
//	}
//	return &pb.HexTx{Tx: hex.EncodeToString(pb.Encode(reply.GetTx()))}, nil
//}
//func testGetHexTxByHashReject(t *testing.T) {
//	var in *pb.ReqHash
//
//	_, err := g.GetHexTxByHash(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetHexTxByHashError(t *testing.T) {
	var in *pb.ReqHash

	qapi.On("QueryTx", in).Return(nil, fmt.Errorf("error")).Once()
	_, err := g.GetHexTxByHash(getOkCtx(), in)
	assert.EqualError(t, err, "error", "the error should be error")
}

func testGetHexTxByHashOK(t *testing.T) {
	var in *pb.ReqHash
	tx := &types.Transaction{Fee: 1}
	var td = &types.TransactionDetail{Tx: tx}
	var tdNil = &types.TransactionDetail{Tx: nil}

	encodetx := hex.EncodeToString(pb.Encode(tx))

	qapi.On("QueryTx", in).Return(tdNil, nil).Once()
	data, err := g.GetHexTxByHash(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, "", data.Tx)

	qapi.On("QueryTx", in).Return(td, nil).Once()
	data, err = g.GetHexTxByHash(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Equal(t, encodetx, data.Tx)

}

func TestGetHexTxByHash(t *testing.T) {
	//testGetHexTxByHashReject(t)
	testGetHexTxByHashError(t)
	testGetHexTxByHashOK(t)
}

//func testGetTransactionByAddrReject(t *testing.T) {
//	var in *pb.ReqAddr
//
//	_, err := g.GetTransactionByAddr(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetTransactionByAddrOK(t *testing.T) {
	var in *pb.ReqAddr
	qapi.On("GetTransactionByAddr", in).Return(nil, nil)
	data, err := g.GetTransactionByAddr(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetTransactionByAddr(t *testing.T) {
	//testGetTransactionByAddrReject(t)
	testGetTransactionByAddrOK(t)
}

//func testGetTransactionByHashesReject(t *testing.T) {
//	var in *pb.ReqHashes
//
//	_, err := g.GetTransactionByHashes(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetTransactionByHashesOK(t *testing.T) {
	var in *pb.ReqHashes
	qapi.On("GetTransactionByHash", in).Return(nil, nil)
	data, err := g.GetTransactionByHashes(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetTransactionByHashes(t *testing.T) {
	//testGetTransactionByHashesReject(t)
	testGetTransactionByHashesOK(t)
}

//func testGetHeadersReject(t *testing.T) {
//	var in *pb.ReqBlocks
//
//	_, err := g.GetHeaders(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetHeadersOK(t *testing.T) {
	var in *pb.ReqBlocks
	qapi.On("GetHeaders", in).Return(nil, nil)
	data, err := g.GetHeaders(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetHeaders(t *testing.T) {
	//testGetHeadersReject(t)
	testGetHeadersOK(t)
}

//func testGetBlockOverviewReject(t *testing.T) {
//	var in *pb.ReqHash
//
//	_, err := g.GetBlockOverview(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetBlockOverviewOK(t *testing.T) {
	var in *pb.ReqHash
	qapi.On("GetBlockOverview", in).Return(nil, nil)
	data, err := g.GetBlockOverview(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetBlockOverview(t *testing.T) {
	//testGetBlockOverviewReject(t)
	testGetBlockOverviewOK(t)
}

//func (g *Grpc) GetAddrOverview(ctx context.Context, in *pb.ReqAddr) (*pb.AddrOverview, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetAddrOverview(in)
//}

//func testGetAddrOverviewReject(t *testing.T) {
//	var in *pb.ReqAddr
//
//	_, err := g.GetAddrOverview(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

//client implement self GetAddrOverview instead of api interface
//func testGetAddrOverviewOK(t *testing.T) {
//	var in *pb.ReqAddr
//	qapi.On("GetAddrOverview", in).Return(nil, nil)
//	data, err := g.GetAddrOverview(getOkCtx(), in)
//	assert.Nil(t, err, "the error should be nil")
//	assert.Nil(t, data)
//
//}

//func TestGetAddrOverview(t *testing.T) {
//	//testGetAddrOverviewReject(t)
//	testGetAddrOverviewOK(t)
//}

//func (g *Grpc) GetBlockHash(ctx context.Context, in *pb.ReqInt) (*pb.ReplyHash, error) {
//	if !g.checkWhitlist(ctx) {
//		return nil, fmt.Errorf("reject")
//	}
//	return g.cli.GetBlockHash(in)
//}
//func testGetBlockHashReject(t *testing.T) {
//	var in *pb.ReqInt
//
//	_, err := g.GetBlockHash(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetBlockHashOK(t *testing.T) {
	var in *pb.ReqInt
	qapi.On("GetBlockHash", in).Return(nil, nil)
	data, err := g.GetBlockHash(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetBlockHash(t *testing.T) {
	//testGetBlockHashReject(t)
	testGetBlockHashOK(t)
}

//func testIsSyncReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.IsSync(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testIsSyncOK(t *testing.T) {
	var in *pb.ReqNil
	qapi.On("IsSync").Return(nil, nil)
	data, err := g.IsSync(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestIsSync(t *testing.T) {
	//testIsSyncReject(t)
	testIsSyncOK(t)
}

//func testIsNtpClockSyncReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.IsNtpClockSync(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testIsNtpClockSyncOK(t *testing.T) {
	var in *pb.ReqNil
	qapi.On("IsNtpClockSync").Return(nil, nil)
	data, err := g.IsNtpClockSync(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestIsNtpClockSync(t *testing.T) {
	//testIsNtpClockSyncReject(t)
	testIsNtpClockSyncOK(t)
}

//func testGetLastHeaderReject(t *testing.T) {
//	var in *pb.ReqNil
//
//	_, err := g.GetLastHeader(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the error should be reject")
//}

func testGetLastHeaderOK(t *testing.T) {
	var in *pb.ReqNil
	qapi.On("GetLastHeader").Return(nil, nil)
	data, err := g.GetLastHeader(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestGetLastHeader(t *testing.T) {
	//testGetLastHeaderReject(t)
	testGetLastHeaderOK(t)
}

//func testCreateRawTransactionReject(t *testing.T) {
//	var in *pb.CreateTx
//
//	_, err := g.CreateRawTransaction(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the erros should be reject")
//}

//func testCreateRawTransactionError(t *testing.T)  {
//	var in *pb.CreateTx
//
//	qapi.On("CreateRawTransaction", in).Return(nil, fmt.Errorf("rejj"))
//	_,err := g.CreateRawTransaction(getOkCtx(), in)
//	assert.EqualError(t,err,"error","return error")
//}

//func testCreateRawTransactionOk(t *testing.T)  {
//	var in *pb.CreateTx
//	reply := []byte("reply")
//
//
//	qapi.On("CreateRawTransaction", in).Return(reply, nil)
//	data,_ := g.CreateRawTransaction(getOkCtx(), in)
//	assert.Equal(t,reply,data.Data,"return correct reply data")
//}
//
//func Test_CreateRawTransaction(t *testing.T) {
//	//testCreateRawTransactionReject(t)
//	testCreateRawTransactionError(t)
//	testCreateRawTransactionOk(t)
//}

//func testSendRawTransactionReject(t *testing.T) {
//	var in *pb.SignedTx
//
//	_, err := g.SendRawTransaction(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the erros should be reject")
//}

//func TestSendRawTransaction(t *testing.T) {
//	testSendRawTransactionReject(t)
//}

//func testQueryTransactionReject(t *testing.T) {
//	var in *pb.ReqHash
//
//	_, err := g.QueryTransaction(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the erros should be reject")
//}
func testQueryTransactionOK(t *testing.T) {
	var in *pb.ReqHash
	qapi.On("QueryTx", in).Return(nil, nil)
	data, err := g.QueryTransaction(getOkCtx(), in)
	assert.Nil(t, err, "the error should be nil")
	assert.Nil(t, data)

}

func TestQueryTransaction(t *testing.T) {
	//testQueryTransactionReject(t)
	testQueryTransactionOK(t)
}

//func testGetBalanceReject(t *testing.T) {
//	var in *pb.ReqBalance
//
//	_, err := g.GetBalance(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the erros should be reject")
//}

//func TestGetBalance(t *testing.T) {
//	testGetBalanceReject(t)
//}

//func testGetTokenBalanceReject(t *testing.T) {
//	var in *pb.ReqTokenBalance
//
//	_, err := g.GetTokenBalance(getNokCtx(), in)
//	assert.EqualError(t, err, "reject", "the erros should be reject")
//}

//func TestGetTokenBalance(t *testing.T) {
//	testGetTokenBalanceReject(t)
//}
