package btcwallet

import (
	"path/filepath"

	"github.com/viacoin/lnd/lnwallet"
	"github.com/roasbeef/btcd/chaincfg"
	"github.com/roasbeef/btcd/wire"
	"github.com/roasbeef/btcutil"

	"github.com/roasbeef/btcwallet/chain"

	// This is required to register bdb as a valid walletdb driver. In the
	// init function of the package, it registers itself. The import is used
	// to activate the side effects w/o actually binding the package name to
	// a file-level variable.
	_ "github.com/roasbeef/btcwallet/walletdb/bdb"
)

var (
	lnwalletHomeDir = btcutil.AppDataDir("lnwallet", false)
	defaultDataDir  = lnwalletHomeDir

	defaultLogFilename = "lnwallet.log"
	defaultLogDirname  = "logs"
	defaultLogDir      = filepath.Join(lnwalletHomeDir, defaultLogDirname)

	btcdHomeDir        = btcutil.AppDataDir("btcd", false)
	btcdHomedirCAFile  = filepath.Join(btcdHomeDir, "rpc.cert")
	defaultRPCKeyFile  = filepath.Join(lnwalletHomeDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(lnwalletHomeDir, "rpc.cert")

	// defaultPubPassphrase is the default public wallet passphrase which is
	// used when the user indicates they do not want additional protection
	// provided by having all public data in the wallet encrypted by a
	// passphrase only known to them.
	defaultPubPassphrase = []byte("public")

	walletDbName = "lnwallet.db"
)

// Config is a struct which houses configuration parameters which modify the
// instance of BtcWallet generated by the New() function.
type Config struct {
	// DataDir is the name of the directory where the wallet's persistent
	// state should be stored.
	DataDir string

	// LogDir is the name of the directory which should be used to store
	// generated log files.
	LogDir string

	// PrivatePass is the private password to the underlying btcwallet
	// instance. Without this, the wallet cannot be decrypted and operated.
	PrivatePass []byte

	// PublicPass is the optional public password to btcwallet. This is
	// optionally used to encrypt public material such as public keys and
	// scripts.
	PublicPass []byte

	// HdSeed is an optional seed to feed into the wallet. If this is
	// unspecified, a new seed will be generated.
	HdSeed []byte

	// ChainSource is the primary chain interface. This is used to operate
	// the wallet and do things such as rescanning, sending transactions,
	// notifications for received funds, etc.
	ChainSource chain.Interface

	// FeeEstimator is an instance of the fee estimator interface which
	// will be used by the wallet to dynamically set transaction fees when
	// crafting transactions.
	FeeEstimator lnwallet.FeeEstimator

	// NetParams is the net parameters for the target chain.
	NetParams *chaincfg.Params
}

// NetworkDir returns the directory name of a network directory to hold wallet
// files.
func NetworkDir(dataDir string, chainParams *chaincfg.Params) string {
	netname := chainParams.Name

	// For now, we must always name the testnet data directory as "testnet"
	// and not "testnet3" or any other version, as the chaincfg testnet3
	// parameters will likely be switched to being named "testnet3" in the
	// future.  This is done to future proof that change, and an upgrade
	// plan to move the testnet3 data directory can be worked out later.
	if chainParams.Net == wire.TestNet3 {
		netname = "testnet"
	}

	return filepath.Join(dataDir, netname)
}
