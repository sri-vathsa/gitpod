// generated using github.com/32leaves/bel on 2021-03-10 16:03:41.509150944 +0000 UTC m=+0.008952353
// DO NOT MODIFY

export enum WorkspaceInitSource {
    WorkspaceInitFromBackup = "from-backup",
    WorkspaceInitFromPrebuild = "from-prebuild",
    WorkspaceInitFromOther = "from-other",
}
export interface WorkspaceReadyMessage {
    source: WorkspaceInitSource
}
