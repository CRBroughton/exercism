pub fn secret_add(secret: Int) -> fn(Int) -> Int {
  fn(add: Int) { add + secret }
}

pub fn secret_subtract(secret: Int) -> fn(Int) -> Int {
  fn(subtract: Int) { subtract - secret }
}

pub fn secret_multiply(secret: Int) -> fn(Int) -> Int {
  fn(multiply: Int) { multiply * secret }
}

pub fn secret_divide(secret: Int) -> fn(Int) -> Int {
  fn(divider: Int) { divider / secret }
}

pub fn secret_combine(
  secret_function1: fn(Int) -> Int,
  secret_function2: fn(Int) -> Int,
) -> fn(Int) -> Int {
  fn(arg: Int) { secret_function2(secret_function1(arg)) }
}
