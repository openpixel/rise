variable "i" {
  value = 6
}

variable "j" {
  value = 100
}

variable "foo" {
  value = ["this", "that"]
}

variable "bar" {
  value = {
    "this" = "that"
  }
}

variable "nested" {
  value = [
    {
      "id" = "1"
    },
    {
      "id" = "2"
    }
  ]
}

template "thing1" {
  content = "the value of j is ${j}"
}
