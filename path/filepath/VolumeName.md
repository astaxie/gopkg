## func VolumeName(path string) (v string)

参数列表

- path: 路径名

返回值列表

- v: 如果是window平台返回分区名,如给定"c:\foo\bar", 将返回"c:", 如果给定"\\host\share\foo", 返回"\\host\share", 其他平台返回""

功能说明

windows平台下返回文件的分区名。
