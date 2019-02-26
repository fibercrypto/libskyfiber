typedef GoSlice_ coin__Transactions;
/**
 * Skycoin transaction.
 *
 * Instances of this struct are included in blocks.
 */
typedef struct {
    GoInt32_ Length;          ///< Current transaction's length expressed in bytes.
    GoInt8_ Type;             ///< Transaction's version. When a node tries to process a transaction, it must verify whether it supports the transaction's type. This is intended to provide a way to update skycoin clients and servers without crashing the network. If the transaction is not compatible with the node, it should not process it.
    cipher__SHA256 InnerHash; ///< It's a SHA256 hash of the inputs and outputs of the transaction. It is used to protect against transaction mutability. This means that the transaction cannot be altered after its creation.

    GoSlice_ Sigs; ///< A list of digital signiatures generated by the skycoin client using the private key. It is used by Skycoin servers to verify the authenticy of the transaction. Each input requires a different signature.
    GoSlice_ In;   ///< A list of references to unspent transaction outputs. Unlike other cryptocurrencies, such as Bitcoin, Skycoin unspent transaction outputs (UX) and Skycoin transactions (TX) are separated in the blockchain protocol, allowing for lighter transactions, thus reducing the broadcasting costs across the network.
    GoSlice_ Out;  ///< Outputs: A list of outputs created by the client, that will be recorded in the blockchain if transactions are confirmed. An output consists of a data structure representing an UTXT, which is composed by a Skycoin address to be sent to, the amount in Skycoin to be sent, and the amount of Coin Hours to be sent, and the SHA256 hash of the previous fields.
} coin__Transaction;

/**
 * Skycoin transaction output.
 *
 * Instances are integral part of transactions included in blocks.
 */
typedef struct {
    cipher__Address Address; ///< Receipient address.
    GoUint64_ Coins;         ///< Amount sent to the receipient address.
    GoUint64_ Hours;         ///< Amount of Coin Hours sent to the receipient address.
} coin__TransactionOutput;
