package main

import "testing"

var tests = []struct {
	root    *TreeNode //the root node of test tree
	kValue  int       //the wanted k-banlanced degree
	expNode string    //the expected node name
}{
	{ //0
		root:    nil,
		kValue:  1,
		expNode: "",
	},
	{ //1
		root: &TreeNode{
			name: "root",
		},
		kValue:  1,
		expNode: "",
	},
	{ //2
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "left",
			},
		},
		kValue:  0,
		expNode: "root",
	},
	{ //3
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "left",
			},
		},
		kValue:  1,
		expNode: "",
	},
	{ //4
		root: &TreeNode{
			name: "root",
			right: &TreeNode{
				name: "right",
			},
		},
		kValue:  0,
		expNode: "root",
	},
	{ //5
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				left: &TreeNode{
					name: "leftB",
				},
			},
		},
		kValue:  1,
		expNode: "root",
	},
	{ //6
		root: &TreeNode{
			name: "root",
			right: &TreeNode{
				name: "right",
				left: &TreeNode{
					name: "left",
				},
			},
		},
		kValue:  1,
		expNode: "root",
	},
	{ //7
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "left",
			},
			right: &TreeNode{
				name: "right",
			},
		},
		kValue:  0,
		expNode: "",
	},
	{ //8
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
			},
			right: &TreeNode{
				name: "right",
				left: &TreeNode{
					name: "leftB",
				},
			},
		},
		kValue:  0,
		expNode: "right",
	},
	{ //9
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				left: &TreeNode{
					name: "leftB",
				},
			},
			right: &TreeNode{
				name: "rightA",
				right: &TreeNode{
					name: "rightB",
				},
			},
		},
		kValue:  0,
		expNode: "leftA",
	},
	{ //10
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				left: &TreeNode{
					name: "leftB",
					right: &TreeNode{
						name: "rightC",
					},
				},
			},
			right: &TreeNode{
				name: "rightA",
				right: &TreeNode{
					name: "rightB",
				},
			},
		},
		kValue:  0,
		expNode: "leftB",
	},
	{ //11
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				left: &TreeNode{
					name: "leftB",
					left: &TreeNode{
						name: "leftC",
						left: &TreeNode{
							name: "leftD",
						},
					},
				},
			},
		},
		kValue:  2,
		expNode: "leftA",
	},
	{ //12
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				left: &TreeNode{
					name: "leftB",
					left: &TreeNode{
						name: "leftC",
						left: &TreeNode{
							name: "leftD",
						},
					},
					right: &TreeNode{
						name: "rightB",
					},
				},
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftE",
					},
				},
			},
		},
		kValue:  2,
		expNode: "root",
	},
	{ //13
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  0,
		expNode: "leftC",
	},
	{ //14
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  1,
		expNode: "rightB",
	},
	{ //15
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  2,
		expNode: "leftB",
	},
	{ //16
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  3,
		expNode: "rightA",
	},
	{ //17
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  4,
		expNode: "leftA",
	},
	{ //18
		root: &TreeNode{
			name: "root",
			left: &TreeNode{
				name: "leftA",
				right: &TreeNode{
					name: "rightA",
					left: &TreeNode{
						name: "leftB",
						right: &TreeNode{
							name: "rightB",
							left: &TreeNode{
								name: "leftC",
								right: &TreeNode{
									name: "rightC",
								},
							},
						},
					},
				},
			},
		},
		kValue:  5,
		expNode: "root",
	},
}

func Test_FindNonKBalancedNode(t *testing.T) {
	for idx, tt := range tests {
		ret := FindNonKBalancedNode(tt.root, tt.kValue)
		if ret.node == nil {
			if tt.expNode != "" {
				t.Errorf("[%d] expected=%s, got=", idx, tt.expNode)
			}
		} else if ret.node.name != tt.expNode {
			t.Errorf("[%d] expected=%s, got=%s",
				idx, tt.expNode, ret.node.name)
		}
	}
}
