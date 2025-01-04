package day17

func (c *Computer) RegisterA() int { return c.registerA }
func (c *Computer) RegisterB() int { return c.registerB }
func (c *Computer) RegisterC() int { return c.registerC }

func (c *Computer) SetRegisterA(in int) { c.registerA = in }
func (c *Computer) SetRegisterB(in int) { c.registerB = in }
func (c *Computer) SetRegisterC(in int) { c.registerC = in }

func (c *Computer) GetData() []uint8   { return c.Program.data }
func (c *Computer) SetData(in []uint8) { c.Program.data = in }
