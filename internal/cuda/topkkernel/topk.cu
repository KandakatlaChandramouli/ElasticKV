extern "C" __global__
void TopKKernel(
    const float* input,
    float* output,
    int k,
    int n
) {
    int idx = threadIdx.x;

    if (idx >= k) return;

    output[idx] = input[idx];
}
