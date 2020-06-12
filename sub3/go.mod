
module "github.com/thaJeztah/monono/sub3"

go 1.13

require (
    github.com/thaJeztah/monono/sub1 v0.0.0-00010101000000-000000000000
    github.com/thaJeztah/monono/sub2 v0.0.0-00010101000000-000000000000
)

replace (
    github.com/thaJeztah/monono/sub1 => ../sub1
    github.com/thaJeztah/monono/sub2 => ../sub2
)
