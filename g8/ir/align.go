package ir

const regSize = 4

func alignUp(size, align int32) int32 {
	mod := size % align
	if mod == 0 {
		return size
	}
	return size + align - mod
}
