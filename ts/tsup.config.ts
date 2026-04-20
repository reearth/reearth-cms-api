import { defineConfig } from "tsup";

export default defineConfig({
  entry: {
    index: "src/index.ts",
    webhook: "src/webhook.ts",
    public: "src/public.ts",
  },
  format: ["esm", "cjs"],
  dts: true,
  sourcemap: true,
  clean: true,
  target: "es2022",
  splitting: false,
  treeshake: true,
});
