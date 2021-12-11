# Read data into list of rows, convert to matrix
f = "test_input.txt"
d = split(strip(read(f, String)), "\n")
l = map(r -> [x - '0' for x in r], d)
m = Matrix(transpose(reshape(vcat(l...), (:, length(l)))))
println("Starting matrix:")
display(m)
nrows, ncols = size(m)

function part1(m)

    # Start the simulation
    totFlashes = 0
    for t in 1:2

        # At the beginning of each step, add 1 to each cell
        m = m .+ 1

        # Any cells with a value > 9 flash, continue until
        # no more flashes
        flashesFound = true
        while flashesFound
            flashesFound = false
            for r in 1:nrows
                for c in 1:ncols
                    if m[r,c] > 9
                        flash(m, r, c)
                        flashesFound = true
                        totFlashes += flashes
                    end
                end
            end
        end

        println("\nAfter step $t:")
        display(m)
        println("Total flashes = $totFlashes")
    end
end

# Flash one cell in the matrix: this sets the cell to zero,
# and increments all diagonal cells by one
function flash(m, r, c)
    for ri in r-1:r+1
        for ci in c-1:c+1
            if r >= 1 && r <= nrows && c >= 1 && c <= ncols
                m[r,c] += 1
            end
        end
    end
    m[r,c] = 0
end        

part1(m)