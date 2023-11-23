package main

// shoing platform product list 
// two capibility
 // 1. injest the particular product in cataloge
 // 2. fiter product base on the pricing  return list product 
 // main maintain the data store inmemory

 // "100-200"---list of Product

 // list 150


 type ProductPrice struct{
	
	Price int
 }



 type Node struct{
	ProductPrice ProductPrice
	Left*Node
	Right*Node
 }

 type BST struct{
	Root*Node
 }

 func(bst*BST)Insert(product ProductPrice){
	if bst.Root==nil{
		bst.Root=&Node{ProductPrice: product,Left: nil,Right: nil}
	}else{
		bst.insertRec(bst.Root,product)
	}

	
 }


 func(bst*BST)insertRec(node*Node,product ProductPrice)*Node{

	if node==nil{
		return &Node{ProductPrice:product,Left: nil,Right: nil}
	}
	if product.Price<node.ProductPrice.Price{
		node.Left=bst.insertRec(node.Left,product)
	
	}else if product.Price>node.ProductPrice.Price{
		node.Right=bst.insertRec(node.Right,product)
	}
	return node

 }

 func(bst*BST)Search(price int)*ProductPrice{

	return bst.searchRec(bst.Root,price)
 }

 func(bst*BST)searchRec(node*Node,price int)*ProductPrice{

	if node==nil||node.ProductPrice.Price==price{
		return&node.ProductPrice
	}

	if price<node.ProductPrice.Price{
		return bst.searchRec(node.Left,price)
	}
	return bst.searchRec(node.Right,price)
 }


 

 func  main(){


	   
	   

	
 }
 