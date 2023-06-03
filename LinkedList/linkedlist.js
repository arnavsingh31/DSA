class Node {
    constructor(data) {
        this.data = data;
        this.next = null;
    }
}

class LinkedList {
    constructor (head = null) {
        this.head = head;
    }

    // Moves last element of the linked list to the front.
    moveLastToFront () {
        let currNode = this.head;
        let lastNode ;
        while (currNode.next != null) {
            if (currNode.next.next == null) {
                lastNode = currNode.next;
                currNode.next = null;
                break;
            }
            currNode = currNode.next;
        }

        lastNode.next = this.head;
        this.head = lastNode;

    }
}

let node1 = new Node(1);
let node2 = new Node(2);
let node3 = new Node(3);
let node4 = new Node(4);
let node5 = new Node(5);
let node6 = new Node(6);

node1.next = node2;
node2.next = node3;
node3.next = node4;
node4.next = node5;
node5.next = node6;

let LL = new LinkedList(node1);

// console.log("before---------->",LL.head.data);
// LL.moveLastToFront();
// console.log("after---------->",LL.head.data);
