extern "C" __global__
void HNSWSearchKernel(
    const float* vectors,
    const float* query,
    float* scores,
    int dim,
    int count
) {

    int idx =
        blockIdx.x *
        blockDim.x +
        threadIdx.x;

    if (idx >= count) return;

    float score = 0.0f;

    int offset = idx * dim;

    for (int i = 0; i < dim; i++) {

        score +=
            vectors[offset + i] *
            query[i];
    }

    scores[idx] = score;
}
