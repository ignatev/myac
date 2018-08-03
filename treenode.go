package main

type treenode struct {
	path, name string
	parent *treenode
	children []*treenode

}

func treelines(tree treenode) {
	var result []string
	result = append(result, tree.name)
	children = 
}


// public class TreeNode<T> implements Iterable<TreeNode<T>> {

//     T data;
//     TreeNode<T> parent;
//     List<TreeNode<T>> children;

//     public TreeNode(T data) {
//         this.data = data;
//         this.children = new LinkedList<TreeNode<T>>();
//     }

//     public TreeNode<T> addChild(T child) {
//         TreeNode<T> childNode = new TreeNode<T>(child);
//         childNode.parent = this;
//         this.children.add(childNode);
//         return childNode;
//     }

//     // other features ...

// }