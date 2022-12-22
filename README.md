# go-vmachine
A Simple Simulation of VM
```go
vmachine := vm.NewVM()
vmachine.Push(
  instruction.NewNumber(20),
  instruction.NewNumber(30),
  instruction.NewOperator(instruction.OpCodeAdd),
  instruction.NewNumber(10),
  instruction.NewOperator(instruction.OpCodeSub),
)
machine_result := vmachine.Run()
fmt.Printf("Result: %v", machine_result)
// Result: OpCode: number | Operand: 40
```
