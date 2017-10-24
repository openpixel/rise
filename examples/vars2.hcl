variable "j" {
  value = "${env("J_VAL")}"
}

variable "k" {
  value = {
    t = "z"
  }
}

variable "h" {
  value = ["Foo"]
}
