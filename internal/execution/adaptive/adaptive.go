package adaptive

func BatchSize(
    rows int,
) int {

    if rows > 100000 {
        return 4096
    }

    if rows > 10000 {
        return 2048
    }

    return 512
}
