#include <stdlib.h>

extern void goRVExtension(char *output, size_t outputSize, char *input);
extern void goRVExtensionVersion(char *output, size_t outputSize);
extern int goRVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc);

#ifdef WIN64
__declspec(dllexport) void RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}
__declspec(dllexport) void RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}
__declspec(dllexport) int RVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc) {
	goRVExtensionArgs(output, outputSize, input, argv, argc);
}
#else
__declspec(dllexport) void __stdcall _RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}
__declspec(dllexport) void __stdcall _RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}
__declspec(dllexport) void __stdcall _RVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc) {
	goRVExtensionArgs(output, outputSize, input, argv, argc);
}
#endif
// do this for all the other exported functions