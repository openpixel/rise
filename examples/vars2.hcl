variable "i" {
  value = 10
}

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
