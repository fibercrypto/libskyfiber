typedef struct{
    GoUint64_ Time;
    GoUint64_ BkSeq;
} coin__UxHead;
typedef GoMap_ coin__UxHashSet;
typedef GoSlice_  coin__UxArray;
typedef GoMap_ coin__AddressUxOuts;
typedef struct{
    coin__UxHead Head;
    coin__UxBody Body;
} coin__UxOut;
typedef struct{
    cipher__SHA256 SrcTransaction;
    cipher__Address Address;
    GoUint64_ Coins;
    GoUint64_ Hours;
} coin__UxBody;