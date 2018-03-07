variable "j" {
    value = "world"
}

template "advanced" {
    file = "./hello.txt"
    trim = true
}
