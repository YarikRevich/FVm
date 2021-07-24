package memory

var (
	Memory = make([]uint16, 0xFFFF)
)

func MemRead(cell uint16)uint16{
	return 0;
};

func MemWrite(src uint16, to uint16){

}