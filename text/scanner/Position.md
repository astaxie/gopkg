	type Position struct {
    	Filename string // 文件名，可能为空
    	Offset   int    // 位移，从0计
    	Line     int    // 行数，从1计
    	Column   int    // 列数，从1计（按字符计算）
	}

Position用来表示源中的位置。行数Line > 0的情况下，Position合法。

