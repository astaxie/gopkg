
```go
// A Decoder manages the receipt of type and data information read from the
// remote side of a connection.
type Decoder struct {
	mutex        sync.Mutex                              // each item must be received atomically
	r            io.Reader                               // source of the data
	buf          bytes.Buffer                            // buffer for more efficient i/o from r
	wireType     map[typeId]*wireType                    // map from remote ID to local description
	decoderCache map[reflect.Type]map[typeId]**decEngine // cache of compiled engines
	ignorerCache map[typeId]**decEngine                  // ditto for ignored objects
	freeList     *decoderState                           // list of free decoderStates; avoids reallocation
	countBuf     []byte                                  // used for decoding integers while parsing messages
	tmp          []byte                                  // temporary storage for i/o; saves reallocating
	err          error
}


// An Encoder manages the transmission of type and data information to the
// other side of a connection.
type Encoder struct {
	mutex      sync.Mutex              // each item must be sent atomically
	w          []io.Writer             // where to send the data
	sent       map[reflect.Type]typeId // which types we've already sent
	countState *encoderState           // stage for writing counts
	freeList   *encoderState           // list of free encoderStates; avoids reallocation
	byteBuf    bytes.Buffer            // buffer for top-level encoderState
	err        error
}
```

