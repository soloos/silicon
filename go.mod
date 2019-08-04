module soloos/soloboat

go 1.12

require (
	soloos/common v0.0.0
	soloos/sdbone v0.0.0
)

replace (
	soloos/common v0.0.0 => /soloos/common
	soloos/sdbone v0.0.0 => /soloos/sdbone
	soloos/sdfs v0.0.0 => /soloos/sdfs
	soloos/soloboat v0.0.0 => /soloos/soloboat
	soloos/swal v0.0.0 => /soloos/swal
)
