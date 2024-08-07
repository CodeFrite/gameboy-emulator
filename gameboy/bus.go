package gameboy

type Bus struct {
	mmu *MMU
}

func NewBus() *Bus {
	return &Bus{
		mmu: NewMMU(),
	}
}

func (b *Bus) AttachMemory(name string, address uint16, memory Accessible) {
	b.mmu.AttachMemory(name, address, memory)
}

func (b *Bus) Read(addr uint16) uint8 {
	return b.mmu.Read(addr)
}

func (b *Bus) Read16(addr uint16) uint16 {
	return b.mmu.Read16(addr)
}

func (b *Bus) Dump(from uint16, to uint16) []uint8 {
	return b.mmu.Dump(from, to)
}

func (b *Bus) Write(addr uint16, value uint8) {
	b.mmu.Write(addr, value)
}

func (b *Bus) WriteBlob(addr uint16, blob []uint8) {
	b.mmu.WriteBlob(addr, blob)
}
