extern "C" __global__
void IVFProbeKernel(
    const int* clusters,
    int* probes,
    int nprobe
) {

    int idx =
        blockIdx.x *
        blockDim.x +
        threadIdx.x;

    if (idx < nprobe) {
        probes[idx] = clusters[idx];
    }
}
