package device

type Device struct {
    ID int
    Name string
    MemoryMB int
}

func Default() Device {

    return Device{
        ID: 0,
        Name: "Tesla T4",
        MemoryMB: 16384,
    }
}
