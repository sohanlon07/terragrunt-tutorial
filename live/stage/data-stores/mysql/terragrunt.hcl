terraform {
    source = "git::github.com/sohanlon07/terraform-tutorial-modules.git//modules/data-stores/mysql?ref=v0.0.25"
}

include {
    
}

inputs = {
    db_name = "example-stage"
}