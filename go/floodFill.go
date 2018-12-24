func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    color :=  image[sr][sc]
    image[sr][sc] = newColor
    if color ==  newColor {
        return image
    }
    
    if sr - 1 >= 0 && color == image[sr-1][sc] {
        floodFill(image, sr - 1, sc, newColor)
    }
    if sr + 1 < len(image) && color == image[sr+1][sc] {
        floodFill(image, sr + 1, sc, newColor)
    }
    if sc - 1 >= 0 && color == image[sr][sc-1] {
        floodFill(image, sr, sc - 1, newColor)
    }
    if sc + 1 < len(image[0]) && color == image[sr][sc+1] {
        floodFill(image, sr, sc + 1, newColor)
    }
    
    return image
}