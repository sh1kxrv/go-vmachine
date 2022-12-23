# go-vmachine

A Simple Simulation of VM

## Example usage

[Code](./cmd/app/app.go)

## Demo

```go
instructions := []*instruction.Instruction{
  instruction.NewNumber(20),
  instruction.NewNumber(30),
  instruction.NewOperator(instruction.OpCodeAdd),
  instruction.NewNumber(10),
  instruction.NewOperator(instruction.OpCodeSub),
}
// -> 40
```

## TODO

- [ ] Read instructions from file & transpile

---

- [ ] Support string instruction
- [ ] Unit tests
- [ ] Solve number casting problem
