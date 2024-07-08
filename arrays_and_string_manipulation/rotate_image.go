package arrays_and_string_manipulation

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
// Approach 1: Rotate Groups of Four Cells
// Observe how the cells move in groups when we rotate the image.
// We can iterate over each group of four cells and rotate them.
func rotateByGroup(matrix [][]int) {
	n := len(matrix[0])

	for i := 0; i < n/2+n%2; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

func rotate(matrix [][]int) {
	transposeMatrix(matrix)
	reflectMatrix(matrix)
}

// transpose aka reverse the matrix around the main diagonal
// e.g.
// [1,2]
// [6,7]
// goes to
// [1,6]
// [2,7]
func transposeMatrix(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

// [1,6]
// [2,7]
// goes to
// [6,1]
// [7,2]
func reflectMatrix(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
}

//class Solution {
//    public void rotate(int[][] matrix) {
//        transpose(matrix);
//        reflect(matrix);
//    }
//
//    public void transpose(int[][] matrix) {
//        int n = matrix.length;
//        for (int i = 0; i < n; i++) {
//            for (int j = i + 1; j < n; j++) {
//                int tmp = matrix[j][i];
//                matrix[j][i] = matrix[i][j];
//                matrix[i][j] = tmp;
//            }
//        }
//    }
//
//    public void reflect(int[][] matrix) {
//        int n = matrix.length;
//        for (int i = 0; i < n; i++) {
//            for (int j = 0; j < n / 2; j++) {
//                int tmp = matrix[i][j];
//                matrix[i][j] = matrix[i][n - j - 1];
//                matrix[i][n - j - 1] = tmp;
//            }
//        }
//    }
//}
