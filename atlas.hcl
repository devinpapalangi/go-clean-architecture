data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/loader"
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:postgres@localhost:5432/go_refreshments?sslmode=disable"  # Define the DB URL here
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}