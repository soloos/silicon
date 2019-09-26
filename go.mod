module soloos/soloboat

go 1.12

replace (
	soloos/common v0.0.0 => /soloos/common
	soloos/soloboat v0.0.0 => /soloos/soloboat
	soloos/solodb v0.0.0 => /soloos/solodb
	soloos/solofs v0.0.0 => /soloos/solofs
	soloos/solomq v0.0.0 => /soloos/solomq
)

require (
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	soloos/common v0.0.0
	soloos/solodb v0.0.0
)
