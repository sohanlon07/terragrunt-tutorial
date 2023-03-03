remote_state {
    backend = "s3"

    generate = {
        path = "backend.tf"
        if_exists = "overwrite"
    }

    config = {
        bucket = "terraform-state-file-storage-sohan-gm"
        key = "${path_relative_to_include()}/terraform.tfstate"
        region = "us-east-2"
        encrypt = true
        dynamodb_table = "terraform-state-file-storage-sohan-locks-gm"
    }
}