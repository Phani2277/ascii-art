echo
go run ./cmd "hello" | cat -e
echo
go run ./cmd "HELLO" | cat -e
echo
go run ./cmd "HeLlo HuMaN" | cat -e
echo
go run ./cmd "1Hello 2There" | cat -e
echo
go run ./cmd "Hello\nThere" | cat -e
echo
go run ./cmd "Hello\n\nThere" | cat -e
echo
go run ./cmd 'hello There 1 to 2!' | cat -e
echo
go run ./cmd "[\]^_ 'a" | cat -e
echo
go run ./cmd "{Hello & There #}" | cat -e
echo
go run ./cmd "MaD3IrA&LiSboN" | cat -e
echo
go run ./cmd "1a\"#FdwHywR&/()=" | cat -e
echo
go run ./cmd "{|}~" | cat -e
echo
go run ./cmd "[\]^_ 'a" | cat -e
echo
go run ./cmd "RGB" | cat -e
echo
go run ./cmd ":;<=>?@" | cat -e
echo
go run ./cmd '\!" #$%&'"'"'()*+,-./' | cat -e
echo
go run ./cmd "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
echo
go run ./cmd "abcdefghijklmnopqrstuvwxyz" | cat -e
echo
go run ./cmd "
dsadsadas
dsadsadsada
dsadsa" | cat -e
