extern "C" __global__
void TopKKernel(
    const float* input,
    float* output,
    int k
) {

    int idx = threadIdx.x;

    if (idx < k) {
        output[idx] = input[idx];
    }
}
