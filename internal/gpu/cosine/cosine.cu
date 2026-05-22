#include <math.h>

extern "C" __global__
void CosineSimilarity(
    const float* a,
    const float* b,
    float* out,
    int dim
) {

    int idx = blockIdx.x * blockDim.x + threadIdx.x;

    float dot = 0.0f;
    float normA = 0.0f;
    float normB = 0.0f;

    int offset = idx * dim;

    for (int i = 0; i < dim; i++) {

        float av = a[offset + i];
        float bv = b[offset + i];

        dot += av * bv;
        normA += av * av;
        normB += bv * bv;
    }

    out[idx] =
        dot /
        (sqrtf(normA) * sqrtf(normB) + 1e-9f);
}
