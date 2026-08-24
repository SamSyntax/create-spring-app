export interface ProjectConfig {
  projectName: string;
  groupId: string;
  packageName: string;
  springBootVersion: string;
  buildTool: "maven" | "gradle" | "gradle-kotlin";
  javaVersion: string;
  dependencies: string[];
}

export interface SpringMetadata {
  dependencies: {
    values: Array<{
      values: Array<{
        id: string;
        name: string;
        versionRange?: string;
      }>;
    }>;
  };
  javaVersion: {
    default: string;
    values: Array<{ id: string; name: string }>;
  };
  bootVersion: {
    default: string;
    values: Array<{ id: string; name: string }>;
  };
}
