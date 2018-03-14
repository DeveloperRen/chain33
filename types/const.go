package types

import (
	"time"
)

var (
	AllowDepositExec       = []string{"ticket"}
	AllowUserExec          = []string{"coins", "ticket", "hashlock", "retrieve", "none", "token", "trade"}
	GenesisAddr            = "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
	GenesisBlockTime int64 = 1514533394
	HotkeyAddr             = "12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
	FundKeyAddr            = "1BQXS6TxaYYG5mADaWij4AxhZZUTpw95a5"
	EmptyValue             = []byte("emptyBVBiCj5jvE15pEiwro8TQRGnJSNsJF") //这字符串表示数据库中的空值
)

const (
	Coin                     int64  = 1e8
	TokenPrecision           int64  = 1e4
	MaxCoin                  int64  = 1e17
	FutureBlockTime          int64  = 16
	CoinReward               int64  = 18 * Coin //用户回报
	CoinDevFund              int64  = 12 * Coin //发展基金回报
	TicketPrice              int64  = 10000 * Coin
	TicketFrozenTime         int64  = 5  //5s only for test
	TicketWithdrawTime       int64  = 10 //10s only for test
	TicketMinerWaitTime      int64  = 2  // 2s only for test
	MinFee                   int64  = 1e5
	MinBalanceTransfer              = 1e6
	MaxTxSize                int64  = 100000   //100K
	MaxBlockSize             int64  = 10000000 //10M
	MaxTxNumber              int64  = 1600     //160
	PowLimitBits             uint32 = uint32(0x1f00ffff)
	TargetTimespan                  = 144 * 16 * time.Second
	TargetTimePerBlock              = 16 * time.Second
	RetargetAdjustmentFactor        = 4
	MaxTxsPerBlock                  = 100000
	TokenNameLenLimit               = 128
	TokenSymbolLenLimit             = 16
	TokenIntroLenLimit              = 1024
	TokenPrecisionLen               = 1e6
	TokenCreatePriceStand           = 10000 * Coin
	InvalidStartTime                = 0
)

const (
	EventTx                   = 1
	EventGetBlocks            = 2
	EventBlocks               = 3
	EventGetBlockHeight       = 4
	EventReplyBlockHeight     = 5
	EventQueryTx              = 6
	EventTransactionDetail    = 7
	EventReply                = 8
	EventTxBroadcast          = 9
	EventPeerInfo             = 10
	EventTxList               = 11
	EventReplyTxList          = 12
	EventAddBlock             = 13
	EventBlockBroadcast       = 14
	EventFetchBlocks          = 15
	EventAddBlocks            = 16
	EventTxHashList           = 17
	EventTxHashListReply      = 18
	EventGetHeaders           = 19
	EventHeaders              = 20
	EventGetMempoolSize       = 21
	EventMempoolSize          = 22
	EventStoreGet             = 23
	EventStoreSet             = 24
	EventStoreGetReply        = 25
	EventStoreSetReply        = 26
	EventReceipts             = 27
	EventExecTxList           = 28
	EventPeerList             = 29
	EventGetLastHeader        = 30
	EventHeader               = 31
	EventAddBlockDetail       = 32
	EventGetMempool           = 33
	EventGetTransactionByAddr = 34
	EventGetTransactionByHash = 35
	EventReplyTxInfo          = 36
	//wallet event
	EventWalletGetAccountList  = 37
	EventWalletAccountList     = 38
	EventNewAccount            = 39
	EventWalletAccount         = 40
	EventWalletTransactionList = 41
	//EventReplyTxList           = 42
	EventWalletImportprivkey = 43
	EventWalletSendToAddress = 44
	EventWalletSetFee        = 45
	EventWalletSetLabel      = 46
	//EventWalletAccount       = 47
	EventWalletMergeBalance = 48
	EventReplyHashes        = 49
	EventWalletSetPasswd    = 50
	EventWalletLock         = 51
	EventWalletUnLock       = 52
	EventTransactionDetails = 53
	EventBroadcastAddBlock  = 54
	EventGetBlockOverview   = 55
	EventGetAddrOverview    = 56
	EventReplyBlockOverview = 57
	EventReplyAddrOverview  = 58
	EventGetBlockHash       = 59
	EventBlockHash          = 60
	EventGetLastMempool     = 61
	EventWalletGetTickets   = 62
	EventMinerStart         = 63
	EventMinerStop          = 64
	EventWalletTickets      = 65
	EventStoreMemSet        = 66
	EventStoreRollback      = 67
	EventStoreCommit        = 68
	EventCheckBlock         = 69
	//seed
	EventGenSeed      = 70
	EventReplyGenSeed = 71
	EventSaveSeed     = 72
	EventGetSeed      = 73
	EventReplyGetSeed = 74
	EventDelBlock     = 75
	//local store
	EventLocalGet            = 76
	EventLocalReplyValue     = 77
	EventLocalList           = 78
	EventLocalSet            = 79
	EventGetWalletStatus     = 80
	EventCheckTx             = 81
	EventReceiptCheckTx      = 82
	EventQuery               = 83
	EventReplyQuery          = 84
	EventFlushTicket         = 85
	EventFetchBlockHeaders   = 86
	EventAddBlockHeaders     = 87
	EventWalletAutoMiner     = 88
	EventReplyWalletStatus   = 89
	EventGetLastBlock        = 90
	EventBlock               = 91
	EventGetTicketCount      = 92
	EventReplyGetTicketCount = 93
	EventDumpPrivkey         = 94
	EventReplyPrivkey        = 95
	EventIsSync              = 96
	EventReplyIsSync         = 97
)

var eventName = map[int]string{
	1:  "EventTx",
	2:  "EventGetBlocks",
	3:  "EventBlocks",
	4:  "EventGetBlockHeight",
	5:  "EventReplyBlockHeight",
	6:  "EventQueryTx",
	7:  "EventTransactionDetail",
	8:  "EventReply",
	9:  "EventTxBroadcast",
	10: "EventPeerInfo",
	11: "EventTxList",
	12: "EventReplyTxList",
	13: "EventAddBlock",
	14: "EventBlockBroadcast",
	15: "EventFetchBlocks",
	16: "EventAddBlocks",
	17: "EventTxHashList",
	18: "EventTxHashListReply",
	19: "EventGetHeaders",
	20: "EventHeaders",
	21: "EventGetMempoolSize",
	22: "EventMempoolSize",
	23: "EventStoreGet",
	24: "EventStoreSet",
	25: "EventStoreGetReply",
	26: "EventStoreSetReply",
	27: "EventReceipts",
	28: "EventExecTxList",
	29: "EventPeerList",
	30: "EventGetLastHeader",
	31: "EventHeader",
	32: "EventAddBlockDetail",
	33: "EventGetMempool",
	34: "EventGetTransactionByAddr",
	35: "EventGetTransactionByHash",
	36: "EventReplyTxInfo",
	37: "EventWalletGetAccountList",
	38: "EventWalletAccountList",
	39: "EventNewAccount",
	40: "EventWalletAccount",
	41: "EventWalletTransactionList",
	//42: "EventReplyTxList",
	43: "EventWalletImportPrivKey",
	44: "EventWalletSendToAddress",
	45: "EventWalletSetFee",
	46: "EventWalletSetLabel",
	//47: "EventWalletAccount",
	48: "EventWalletMergeBalance",
	49: "EventReplyHashes",
	50: "EventWalletSetPasswd",
	51: "EventWalletLock",
	52: "EventWalletUnLock",
	53: "EventTransactionDetails",
	54: "EventBroadcastAddBlock",
	55: "EventGetBlockOverview",
	56: "EventGetAddrOverview",
	57: "EventReplyBlockOverview",
	58: "EventReplyAddrOverview",
	59: "EventGetBlockHash",
	60: "EventBlockHash",
	61: "EventGetLastMempool",
	62: "EventWalletGetTickets",
	63: "EventMinerStart",
	64: "EventMinerStop",
	65: "EventWalletTickets",
	66: "EventStoreMemSet",
	67: "EventStoreRollback",
	68: "EventStoreCommit",
	69: "EventCheckBlock",
	70: "EventGenSeed",
	71: "EventReplyGenSeed",
	72: "EventSaveSeed",
	73: "EventGetSeed",
	74: "EventReplyGetSeed",
	75: "EventDelBlock",
	76: "EventLocalGet",
	77: "EventLocalReplyValue",
	78: "EventLocalList",
	79: "EventLocalSet",
	80: "EventGetWalletStatus",
	81: "EventCheckTx",
	82: "EventReceiptCheckTx",
	83: "EventQuery",
	84: "EventReplyQuery",
	85: "EventFlushTicket",
	86: "EventFetchBlockHeaders",
	87: "EventAddBlockHeaders",
	88: "EventWalletAutoMiner",
	89: "EventReplyWalletStatus",
	90: "EventGetLastBlock",
	91: "EventBlock",
	92: "EventGetTicketCount",
	93: "EventReplyGetTicketCount",
	94: "EventDumpPrivkey",
	95: "EventReplyPrivkey",
	96: "EventIsSync",
	97: "EventReplyIsSync",
}

//ty = 1 -> secp256k1
//ty = 2 -> ed25519
//ty = 3 -> sm2
const (
	SECP256K1 = 1
	ED25519   = 2
	SM2       = 3
)

//log type
const (
	TyLogErr = 1
	TyLogFee = 2
	//coins
	TyLogTransfer        = 3
	TyLogGenesis         = 4
	TyLogDeposit         = 5
	TyLogExecTransfer    = 6
	TyLogExecWithdraw    = 7
	TyLogExecDeposit     = 8
	TyLogExecFrozen      = 9
	TyLogExecActive      = 10
	TyLogGenesisTransfer = 11
	TyLogGenesisDeposit  = 12

	//log for ticket
	TyLogNewTicket   = 111
	TyLogCloseTicket = 112
	TyLogMinerTicket = 113
	TyLogTicketBind  = 114

	//log for token create
	TyLogPreCreateToken    = 211
	TyLogFinishCreateToken = 212
	TyLogRevokeCreateToken = 213

	//log for trade
	TyLogTradeSell         = 311
	TyLogTradeBuy          = 312
	TyLogTradeRevoke       = 313
)

//exec type
const (
	ExecErr  = 0
	ExecPack = 1
	ExecOk   = 2
)

//coinsaction
const (
	CoinsActionTransfer = 1
	CoinsActionGenesis  = 2
	CoinsActionWithdraw = 3
)

//ticket
const (
	TicketActionGenesis = 11
	TicketActionOpen    = 12
	TicketActionClose   = 13
	TicketActionList    = 14 //读的接口不直接经过transaction
	TicketActionInfos   = 15 //读的接口不直接经过transaction
	TicketActionMiner   = 16
	TicketActionBind    = 17
)

//hashlock const
const (
	HashlockActionLock   = 1
	HashlockActionSend   = 2
	HashlockActionUnlock = 3
)

//retrieve
const (
	RetrievePre    = 1
	RetrievePerf   = 2
	RetrieveBackup = 3
	RetrieveCancel = 4
)

const (
	ActionTransfer = iota
	ActionGenesis
	ActionWithdraw
    TokenActionPreCreate
	TokenActionFinishCreate
	TokenActionRevokeCreate
)

const (
	TokenStatusPreCreated    = iota
	TokenStatusCreated
	TokenStatusCreateRevoked
)

const (
	TradeSell = iota
	TradeBuy
	TradeRevokeSell
)

//0->not start, 1->on sale, 2->sold out, 3->revoke, 4->expired
const (
	NotStart = iota
	OnSale
	SoldOut
	Revoked
	Expired
)
