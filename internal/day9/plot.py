import matplotlib.pyplot as plt
import numpy as np

def plot_points(points, close_path=False, figsize=(6,6), title="Point path"):
    """
    Plot a list of 2D points and connect consecutive points with edges.
    
    Args:
        points: list of (x, y) pairs (tuples or lists) or an Nx2 numpy array.
        close_path: if True, connect the last point back to the first.
        annotate: if True, label each point with its index.
        figsize: tuple with figure size (width, height).
        title: plot title.
    """
    pts = np.asarray(points, dtype=float)
    if pts.ndim != 2 or pts.shape[1] != 2:
        raise ValueError("points must be an iterable of (x, y) pairs")
    
    xs = pts[:, 0]
    ys = pts[:, 1]
    
    # If closing the polygon/path, append first point to the end for plotting
    if close_path:
        xs = np.append(xs, xs[0])
        ys = np.append(ys, ys[0])
    
    fig, ax = plt.subplots(figsize=figsize)
    ax.plot(xs, ys, marker='.', markersize=1)  # don't explicitly set color (uses matplotlib default)
    ax.set_aspect('equal', 'box')
    ax.grid(True, linestyle='--', linewidth=0.5)
    ax.set_xlabel('x')
    ax.set_ylabel('y')
    ax.set_title(title)
    plt.savefig("plot.png", dpi=200)
    return fig, ax

with open("./inputs/day9.txt", "r") as f:
    coords = [[int(n) for n in line.strip().split(",")] for line in f]

# Also show a closed path example
plot_points(coords, close_path=True, title="Example path (closed)")
