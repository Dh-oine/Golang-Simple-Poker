package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func duplicateInArray(arr map[int]int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}
func main() {
	suit := [4]string{"diamonds", "hearts", "clubs", "spades"}
	value := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	cards := make(map[int]string)
	cardsrand := make(map[int]string)
	cardsrandsuits := make(map[int]string)
	cardsrandvalues := make(map[int]string)
	randcardint := make(map[int]int)
	rcsuits := make(map[int]string)
	rcvalues := make(map[int]string)
	p1power := 0
	p2power := 0
	p1kicker := 0
	p2kicker := 0
	sus := 0
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 13; j++ {
		for s := 0; s < 4; s++ {
			cards[sus] = value[j] + " of " + suit[s]
			sus++
		}
	}
	for i := 0; i < 9; i++ {
		randcardint[i] = rand.Intn(len(cards))
		if duplicateInArray(randcardint) > -1 {
			randcardint[i] = rand.Intn(len(cards))
			i--
		}
	}
	for i := 0; i < 9; i++ {
		cardsrand[i] = cards[randcardint[i]]
		cardsrandsuits[i] = cards[randcardint[i]]
		cardsrandvalues[i] = cards[randcardint[i]]
		fmt.Println(cardsrand[i])
	}
	for i := 0; i < 9; i++ {
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "A of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "K of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "Q of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "J of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "10 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "9 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "8 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "7 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "6 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "5 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "4 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "3 of ", "")
		cardsrandsuits[i] = strings.ReplaceAll(cardsrandsuits[i], "2 of ", "")
		rcsuits[i] = cardsrandsuits[i]
		fmt.Println(rcsuits[i])
	}
	for i := 0; i < 9; i++ {
		cardsrandvalues[i] = strings.ReplaceAll(cardsrandvalues[i], " of hearts", "")
		cardsrandvalues[i] = strings.ReplaceAll(cardsrandvalues[i], " of diamonds", "")
		cardsrandvalues[i] = strings.ReplaceAll(cardsrandvalues[i], " of spades", "")
		cardsrandvalues[i] = strings.ReplaceAll(cardsrandvalues[i], " of clubs", "")

		rcvalues[i] = cardsrandvalues[i]
		fmt.Println(rcvalues[i])
	}
	p1cv1 := 0
	p1cv2 := 0
	tc1v := 0
	tc2v := 0
	tc3v := 0
	tc4v := 0
	tc5v := 0
	p2cv1 := 0
	p2cv2 := 0
	for i := 0; i < len(value); i++ {
		if rcvalues[0] == value[i] {
			p1cv1 = i
		}
		if rcvalues[1] == value[i] {
			p1cv2 = i
		}
		if rcvalues[2] == value[i] {
			tc1v = i
		}
		if rcvalues[3] == value[i] {
			tc2v = i
		}
		if rcvalues[4] == value[i] {
			tc3v = i
		}
		if rcvalues[5] == value[i] {
			tc4v = i
		}
		if rcvalues[6] == value[i] {
			tc5v = i
		}
		if rcvalues[7] == value[i] {
			p2cv1 = i
		}
		if rcvalues[8] == value[i] {
			p2cv2 = i
		}
	}
	flash := []string{rcsuits[0], rcsuits[1], rcsuits[2], rcsuits[3], rcsuits[4], rcsuits[5], rcsuits[6]}
	sort.Strings(flash)
	strit := []int{p1cv1, p1cv2, tc1v, tc2v, tc3v, tc4v, tc5v}
	sort.Ints(strit)
	strit2 := make(map[int]int)
	strit3 := make(map[int]int)
	j := 0
	for i := 0; i < 7; i++ {
		if i == 6 {
			strit2[j] = strit[i]
			break
		}
		if strit[i] != strit[i+1] {

			strit2[j] = strit[i]
			j++
		}
	}
	stritp2 := []int{p2cv1, p2cv2, tc1v, tc2v, tc3v, tc4v, tc5v}
	sort.Ints(stritp2)
	for i := 0; i < 7; i++ {
		if i == 6 {
			strit3[j] = stritp2[i]
			break
		}
		if stritp2[i] != stritp2[i+1] {

			strit3[j] = stritp2[i]
			j++
		}
	}
	for i := 0; i < 2; i++ {
		if (len(strit2) == 7) && ((strit2[6] == 12) && (strit2[5] == 11) && (strit2[4] == 10) && (strit2[3] == 9) && (strit2[2] == 8)) && ((flash[6] == flash[5]) && (flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2])) {
			fmt.Println("Royal straight flash of " + flash[6])
			if i == 0 {
				p1power = 1000
			} else {
				p2power = 1000
			}
		} else {
			if (len(strit2) == 7) && ((strit2[5] == 12) && (strit2[4] == 11) && (strit2[3] == 10) && (strit2[2] == 9) && (strit2[1] == 8)) && ((flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1])) {
				fmt.Println("Royal straight flash of " + flash[5])
				if i == 0 {
					p1power = 1000
				} else {
					p2power = 1000
				}
			} else {
				if (len(strit2) == 7) && ((strit2[4] == 12) && (strit2[3] == 11) && (strit2[2] == 10) && (strit2[1] == 9) && (strit2[0] == 8)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
					fmt.Println("Royal straight flash of " + flash[4])
					if i == 0 {
						p1power = 1000
					} else {
						p2power = 1000
					}
				} else {
					if (len(strit2) == 6) && ((strit2[5] == 12) && (strit2[4] == 11) && (strit2[3] == 10) && (strit2[2] == 9) && (strit2[1] == 8)) && ((flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1])) {
						fmt.Println("Royal straight flash of " + flash[5])
						if i == 0 {
							p1power = 1000
						} else {
							p2power = 1000
						}
					} else {
						if (len(strit2) == 6) && ((strit2[4] == 12) && (strit2[3] == 11) && (strit2[2] == 10) && (strit2[1] == 9) && (strit2[0] == 8)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
							fmt.Println("Royal straight flash of " + flash[4])
							if i == 0 {
								p1power = 1000
							} else {
								p2power = 1000
							}
						} else {
							if (len(strit2) == 5) && ((strit2[4] == 12) && (strit2[3] == 11) && (strit2[2] == 10) && (strit2[1] == 9) && (strit2[0] == 8)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
								fmt.Println("Royal straight flash of " + flash[4])
								if i == 0 {
									p1power = 1000
								} else {
									p2power = 1000
								}
							} else {
								if (len(strit2) == 7) && ((strit2[6] == strit2[5]+1) && (strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1)) && ((flash[6] == flash[5]) && (flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2])) {
									fmt.Println("Straight flash from " + value[strit2[2]] + " to " + value[strit2[6]] + " of " + flash[6])
									if i == 0 {
										p1power = strit[6] + 400
									} else {
										p2power = strit[6] + 400
									}
								} else {
									if (len(strit2) == 7) && ((strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1)) && ((flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1])) {
										fmt.Println("Straight flash from " + value[strit2[1]] + " to " + value[strit2[5]] + " of " + flash[5])
										if i == 0 {
											p1power = strit[5] + 400
										} else {
											p2power = strit[5] + 400
										}
									} else {
										if (len(strit2) == 7) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
											fmt.Println("Straight flash from " + value[strit2[0]] + " to " + value[strit2[4]] + " of " + flash[4])
											if i == 0 {
												p1power = strit[4] + 400
											} else {
												p2power = strit[4] + 400
											}
										} else {
											if (len(strit2) == 6) && ((strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1)) && ((flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1])) {
												fmt.Println("Straight flash from " + value[strit2[1]] + " to " + value[strit2[5]] + " of " + flash[5])
												if i == 0 {
													p1power = strit[5] + 400
												} else {
													p2power = strit[5] + 400
												}
											} else {
												if (len(strit2) == 6) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
													fmt.Println("Straight flash from " + value[strit2[0]] + " to " + value[strit2[4]] + " of " + flash[4])
													if i == 0 {
														p1power = strit[4] + 400
													} else {
														p2power = strit[4] + 400
													}
												} else {
													if (len(strit2) == 5) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) && ((flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0])) {
														fmt.Println("Straight flash from " + value[strit2[0]] + " to " + value[strit2[4]] + " of " + flash[4])
														if i == 0 {
															p1power = strit[4] + 400
														} else {
															p2power = strit[4] + 400
														}
													} else {
														if (strit[6] == strit[5]) && (strit[5] == strit[4]) && (strit[4] == strit[3]) {
															fmt.Println("Four of kind " + value[strit[6]])
															if i == 0 {
																p1power = strit[6] + 300
																p1kicker = strit[2]
															} else {
																p2power = strit[6] + 300
																p2kicker = strit[2]
															}
														} else {
															if (strit[5] == strit[4]) && (strit[4] == strit[3]) && (strit[3] == strit[2]) {
																fmt.Println("Four of kind " + value[strit[5]])
																if i == 0 {
																	p1power = strit[5] + 300
																	p1kicker = strit[6]
																} else {
																	p2power = strit[5] + 300
																	p2kicker = strit[6]
																}
															} else {
																if (strit[4] == strit[3]) && (strit[3] == strit[2]) && (strit[2] == strit[1]) {
																	fmt.Println("Four of kind " + value[strit[4]])
																	if i == 0 {
																		p1power = strit[4] + 300
																		p1kicker = strit[6]
																	} else {
																		p2power = strit[4] + 300
																		p2kicker = strit[6]
																	}
																} else {
																	if (strit[3] == strit[2]) && (strit[2] == strit[1]) && (strit[1] == strit[0]) {
																		fmt.Println("Four of kind " + value[strit[3]])
																		if i == 0 {
																			p1power = strit[3] + 300
																			p1kicker = strit[6]
																		} else {
																			p2power = strit[3] + 300
																			p2kicker = strit[6]
																		}
																	} else {
																		if (strit[6] == strit[5]) && (strit[5] == strit[4]) && (strit[3] == strit[2]) {
																			fmt.Println("Full house of " + value[strit[6]] + " and " + value[strit[3]])
																			if i == 0 {
																				p1power = strit[6] + strit[6] + strit[3] + 200
																			} else {
																				p2power = strit[6] + strit[6] + strit[3] + 200
																			}
																		} else {
																			if (strit[6] == strit[5]) && (strit[5] == strit[4]) && (strit[2] == strit[1]) {
																				fmt.Println("Full house of " + value[strit[6]] + " and " + value[strit[2]])
																				if i == 0 {
																					p1power = strit[6] + strit[6] + strit[2] + 200
																				} else {
																					p2power = strit[6] + strit[6] + strit[2] + 200
																				}
																			} else {
																				if (strit[6] == strit[5]) && (strit[5] == strit[4]) && (strit[1] == strit[0]) {
																					fmt.Println("Full house of " + value[strit[6]] + " and " + value[strit[1]])
																					if i == 0 {
																						p1power = strit[6] + strit[6] + strit[1] + 200
																					} else {
																						p2power = strit[6] + strit[6] + strit[1] + 200
																					}
																				} else {
																					if (strit[5] == strit[4]) && (strit[4] == strit[3]) && (strit[2] == strit[1]) {
																						fmt.Println("Full house of " + value[strit[5]] + " and " + value[strit[2]])
																						if i == 0 {
																							p1power = strit[5] + strit[5] + strit[2] + 200
																						} else {
																							p2power = strit[5] + strit[5] + strit[2] + 200
																						}
																					} else {
																						if (strit[5] == strit[4]) && (strit[4] == strit[3]) && (strit[1] == strit[0]) {
																							fmt.Println("Full house of " + value[strit[5]] + " and " + value[strit[1]])
																							if i == 0 {
																								p1power = strit[5] + strit[5] + strit[1] + 200
																							} else {
																								p2power = strit[5] + strit[5] + strit[1] + 200
																							}
																						} else {
																							if (strit[4] == strit[3]) && (strit[3] == strit[2]) && (strit[1] == strit[0]) {
																								fmt.Println("Full house of " + value[strit[4]] + " and " + value[strit[1]])
																								if i == 0 {
																									p1power = strit[4] + strit[4] + strit[1] + 200
																								} else {
																									p2power = strit[4] + strit[4] + strit[1] + 200
																								}
																							} else {
																								if (strit[6] == strit[5]) && (strit[4] == strit[3]) && (strit[3] == strit[2]) {
																									fmt.Println("Full house of " + value[strit[4]] + " and " + value[strit[6]])
																									if i == 0 {
																										p1power = strit[4] + strit[4] + strit[6] + 200
																									} else {
																										p2power = strit[4] + strit[4] + strit[6] + 200
																									}
																								} else {
																									if (strit[6] == strit[5]) && (strit[3] == strit[2]) && (strit[2] == strit[1]) {
																										fmt.Println("Full house of " + value[strit[3]] + " and " + value[strit[6]])
																										if i == 0 {
																											p1power = strit[3] + strit[3] + strit[6] + 200
																										} else {
																											p2power = strit[3] + strit[3] + strit[6] + 200
																										}
																									} else {
																										if (strit[6] == strit[5]) && (strit[2] == strit[1]) && (strit[1] == strit[0]) {
																											fmt.Println("Full house of " + value[strit[2]] + " and " + value[strit[6]])
																											if i == 0 {
																												p1power = strit[2] + strit[2] + strit[6] + 200
																											} else {
																												p2power = strit[2] + strit[2] + strit[6] + 200
																											}
																										} else {
																											if (strit[5] == strit[4]) && (strit[3] == strit[2]) && (strit[2] == strit[1]) {
																												fmt.Println("Full house of " + value[strit[3]] + " and " + value[strit[5]])
																												if i == 0 {
																													p1power = strit[3] + strit[3] + strit[5] + 200
																												} else {
																													p2power = strit[3] + strit[3] + strit[5] + 200
																												}
																											} else {
																												if (strit[5] == strit[4]) && (strit[2] == strit[1]) && (strit[1] == strit[0]) {
																													fmt.Println("Full house of " + value[strit[2]] + " and " + value[strit[5]])
																													if i == 0 {
																														p1power = strit[2] + strit[2] + strit[5] + 200
																													} else {
																														p2power = strit[2] + strit[2] + strit[5] + 200
																													}
																												} else {
																													if (strit[4] == strit[3]) && (strit[2] == strit[1]) && (strit[1] == strit[0]) {
																														fmt.Println("Full house of " + value[strit[2]] + " and " + value[strit[4]])
																														if i == 0 {
																															p1power = strit[2] + strit[2] + strit[4] + 200
																														} else {
																															p2power = strit[2] + strit[2] + strit[4] + 200
																														}
																													} else {
																														if (flash[6] == flash[5]) && (flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) {
																															fmt.Println("Flash of " + flash[6])
																															if i == 0 {
																																p1power = 200
																															} else {
																																p2power = 200
																															}
																														} else {
																															if (flash[5] == flash[4]) && (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) {
																																fmt.Println("Flash of " + flash[5])
																																if i == 0 {
																																	p1power = 200
																																} else {
																																	p2power = 200
																																}
																															} else {
																																if (flash[4] == flash[3]) && (flash[3] == flash[2]) && (flash[2] == flash[1]) && (flash[1] == flash[0]) {
																																	fmt.Println("Flash of " + flash[4])
																																	if i == 0 {
																																		p1power = 200
																																	} else {
																																		p2power = 200
																																	}
																																} else {
																																	if (len(strit2) == 7) && ((strit2[6] == strit2[5]+1) && (strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1)) {
																																		fmt.Println("Straight from " + value[strit2[2]] + " to " + value[strit2[6]])
																																		if i == 0 {
																																			p1power = strit[6] + 100
																																		} else {
																																			p2power = strit[6] + 100
																																		}
																																	} else {
																																		if (len(strit2) == 7) && ((strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1)) {
																																			fmt.Println("Straight from " + value[strit2[1]] + " to " + value[strit2[5]])
																																			if i == 0 {
																																				p1power = strit[5] + 100
																																			} else {
																																				p2power = strit[5] + 100
																																			}
																																		} else {
																																			if (len(strit2) == 7) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) {
																																				fmt.Println("Straight from " + value[strit2[0]] + " to " + value[strit2[4]])
																																				if i == 0 {
																																					p1power = strit[4] + 100
																																				} else {
																																					p2power = strit[4] + 100
																																				}
																																			} else {
																																				if (len(strit2) == 6) && ((strit2[5] == strit2[4]+1) && (strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1)) {
																																					fmt.Println("Straight from " + value[strit2[1]] + " to " + value[strit2[5]])
																																					if i == 0 {
																																						p1power = strit[5] + 100
																																					} else {
																																						p2power = strit[5] + 100
																																					}
																																				} else {
																																					if (len(strit2) == 6) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) {
																																						fmt.Println("Straight from " + value[strit2[0]] + " to " + value[strit2[4]])
																																						if i == 0 {
																																							p1power = strit[4] + 100
																																						} else {
																																							p2power = strit[4] + 100
																																						}
																																					} else {
																																						if (len(strit2) == 5) && ((strit2[4] == strit2[3]+1) && (strit2[3] == strit2[2]+1) && (strit2[2] == strit2[1]+1) && (strit2[1] == strit2[0]+1)) {
																																							fmt.Println("Straight from " + value[strit2[0]] + " to " + value[strit2[4]])
																																							if i == 0 {
																																								p1power = strit[4] + 100
																																							} else {
																																								p2power = strit[4] + 100
																																							}
																																						} else {
																																							if (strit[6] == strit[5]) && (strit[5] == strit[4]) {
																																								fmt.Println("Three of kind " + value[strit[6]])
																																								if i == 0 {
																																									p1power = strit[6] + 80
																																									p1kicker = strit[3]
																																								} else {
																																									p2power = strit[6] + 80
																																									p2kicker = strit[3]
																																								}
																																							} else {
																																								if (strit[5] == strit[4]) && (strit[4] == strit[3]) {
																																									fmt.Println("Three of kind " + value[strit[5]])
																																									if i == 0 {
																																										p1power = strit[5] + 80
																																										p1kicker = strit[6]
																																									} else {
																																										p2power = strit[5] + 80
																																										p2kicker = strit[6]
																																									}
																																								} else {
																																									if (strit[4] == strit[3]) && (strit[3] == strit[2]) {
																																										fmt.Println("Three of kind " + value[strit[4]])
																																										if i == 0 {
																																											p1power = strit[4] + 80
																																											p1kicker = strit[6]
																																										} else {
																																											p2power = strit[4] + 80
																																											p2kicker = strit[6]
																																										}
																																									} else {
																																										if (strit[3] == strit[2]) && (strit[2] == strit[1]) {
																																											fmt.Println("Three of kind " + value[strit[3]])
																																											if i == 0 {
																																												p1power = strit[3] + 80
																																												p1kicker = strit[6]
																																											} else {
																																												p2power = strit[3] + 80
																																												p2kicker = strit[6]
																																											}
																																										} else {
																																											if (strit[2] == strit[1]) && (strit[1] == strit[0]) {
																																												fmt.Println("Three of kind " + value[strit[2]])
																																												if i == 0 {
																																													p1power = strit[2] + 80
																																													p1kicker = strit[6]
																																												} else {
																																													p2power = strit[2] + 80
																																													p2kicker = strit[6]
																																												}
																																											} else {
																																												if (strit[6] == strit[5]) && (strit[4] == strit[3]) {
																																													fmt.Println("Two pairs of " + value[strit[6]] + " and " + value[strit[4]])
																																													if i == 0 {
																																														p1power = strit[6] + strit[4] + 40
																																														p1kicker = strit[2]
																																													} else {
																																														p2power = strit[6] + strit[4] + 40
																																														p2kicker = strit[2]
																																													}
																																												} else {
																																													if (strit[6] == strit[5]) && (strit[3] == strit[2]) {
																																														fmt.Println("Two pairs of " + value[strit[6]] + " and " + value[strit[3]])
																																														if i == 0 {
																																															p1power = strit[6] + strit[3] + 40
																																															p1kicker = strit[4]
																																														} else {
																																															p2power = strit[6] + strit[3] + 40
																																															p2kicker = strit[4]
																																														}
																																													} else {
																																														if (strit[6] == strit[5]) && (strit[2] == strit[1]) {
																																															fmt.Println("Two pairs of " + value[strit[6]] + " and " + value[strit[2]])
																																															if i == 0 {
																																																p1power = strit[6] + strit[2] + 40
																																																p1kicker = strit[4]
																																															} else {
																																																p2power = strit[6] + strit[2] + 40
																																																p2kicker = strit[4]
																																															}
																																														} else {
																																															if (strit[6] == strit[5]) && (strit[1] == strit[0]) {
																																																fmt.Println("Two pairs of " + value[strit[6]] + " and " + value[strit[1]])
																																																if i == 0 {
																																																	p1power = strit[6] + strit[1] + 40
																																																	p1kicker = strit[4]
																																																} else {
																																																	p2power = strit[6] + strit[1] + 40
																																																	p2kicker = strit[4]
																																																}
																																															} else {
																																																if (strit[5] == strit[4]) && (strit[3] == strit[2]) {
																																																	fmt.Println("Two pairs of " + value[strit[5]] + " and " + value[strit[3]])
																																																	if i == 0 {
																																																		p1power = strit[5] + strit[3] + 40
																																																		p1kicker = strit[6]
																																																	} else {
																																																		p2power = strit[5] + strit[3] + 40
																																																		p2kicker = strit[6]
																																																	}
																																																} else {
																																																	if (strit[5] == strit[4]) && (strit[2] == strit[1]) {
																																																		fmt.Println("Two pairs of " + value[strit[5]] + " and " + value[strit[2]])
																																																		if i == 0 {
																																																			p1power = strit[5] + strit[2] + 40
																																																			p1kicker = strit[6]
																																																		} else {
																																																			p2power = strit[5] + strit[2] + 40
																																																			p2kicker = strit[6]
																																																		}
																																																	} else {
																																																		if (strit[5] == strit[4]) && (strit[1] == strit[0]) {
																																																			fmt.Println("Two pairs of " + value[strit[5]] + " and " + value[strit[1]])
																																																			if i == 0 {
																																																				p1power = strit[5] + strit[1] + 40
																																																				p1kicker = strit[6]
																																																			} else {
																																																				p2power = strit[5] + strit[1] + 40
																																																				p2kicker = strit[6]
																																																			}
																																																		} else {
																																																			if (strit[4] == strit[3]) && (strit[2] == strit[1]) {
																																																				fmt.Println("Two pairs of " + value[strit[4]] + " and " + value[strit[2]])
																																																				if i == 0 {
																																																					p1power = strit[4] + strit[2] + 40
																																																					p1kicker = strit[6]
																																																				} else {
																																																					p2power = strit[4] + strit[2] + 40
																																																					p2kicker = strit[6]
																																																				}
																																																			} else {
																																																				if (strit[4] == strit[3]) && (strit[1] == strit[0]) {
																																																					fmt.Println("Two pairs of " + value[strit[4]] + " and " + value[strit[1]])
																																																					if i == 0 {
																																																						p1power = strit[4] + strit[1] + 40
																																																						p1kicker = strit[6]
																																																					} else {
																																																						p2power = strit[4] + strit[1] + 40
																																																						p2kicker = strit[6]
																																																					}
																																																				} else {
																																																					if (strit[3] == strit[2]) && (strit[1] == strit[0]) {
																																																						fmt.Println("Two pairs of " + value[strit[3]] + " and " + value[strit[1]])
																																																						if i == 0 {
																																																							p1power = strit[3] + strit[1] + 40
																																																							p1kicker = strit[6]
																																																						} else {
																																																							p2power = strit[3] + strit[1] + 40
																																																							p2kicker = strit[6]
																																																						}
																																																					} else {
																																																						if (strit[6] == strit[5]) || (strit[6] == strit[4]) || (strit[6] == strit[3]) || (strit[6] == strit[2]) || (strit[6] == strit[1]) || (strit[6] == strit[0]) {
																																																							fmt.Println("Pair of " + value[strit[6]])
																																																							if i == 0 {
																																																								p1power = strit[6] + 20
																																																								p1kicker = strit[5]
																																																							} else {
																																																								p2power = strit[6] + 20
																																																								p2kicker = strit[5]
																																																							}
																																																						} else {
																																																							if (strit[5] == strit[4]) || (strit[5] == strit[3]) || (strit[5] == strit[2]) || (strit[5] == strit[1]) || (strit[5] == strit[0]) {
																																																								fmt.Println("Pair of " + value[strit[5]])
																																																								if i == 0 {
																																																									p1power = strit[5] + 20
																																																									p1kicker = strit[6]
																																																								} else {
																																																									p2power = strit[5] + 20
																																																									p2kicker = strit[6]
																																																								}
																																																							} else {
																																																								if (strit[4] == strit[3]) || (strit[4] == strit[2]) || (strit[4] == strit[1]) || (strit[4] == strit[0]) {
																																																									fmt.Println("Pair of " + value[strit[4]])
																																																									if i == 0 {
																																																										p1power = strit[4] + 20
																																																										p1kicker = strit[6]
																																																									} else {
																																																										p2power = strit[4] + 20
																																																										p2kicker = strit[6]
																																																									}
																																																								} else {
																																																									if (strit[3] == strit[2]) || (strit[3] == strit[1]) || (strit[3] == strit[0]) {
																																																										fmt.Println("Pair of " + value[strit[3]])
																																																										if i == 0 {
																																																											p1power = strit[3] + 20
																																																											p1kicker = strit[6]
																																																										} else {
																																																											p2power = strit[3] + 20
																																																											p2kicker = strit[6]
																																																										}
																																																									} else {
																																																										if (strit[2] == strit[1]) || (strit[2] == strit[0]) {
																																																											fmt.Println("Pair of " + value[strit[2]])
																																																											if i == 0 {
																																																												p1power = strit[2] + 20
																																																												p1kicker = strit[6]
																																																											} else {
																																																												p2power = strit[2] + 20
																																																												p2kicker = strit[6]
																																																											}
																																																										} else {
																																																											if strit[1] == strit[0] {
																																																												fmt.Println("Pair of " + value[strit[1]])
																																																												if i == 0 {
																																																													p1power = strit[1] + 20
																																																													p1kicker = strit[6]
																																																												} else {
																																																													p2power = strit[1] + 20
																																																													p2kicker = strit[6]
																																																												}
																																																											} else {
																																																												fmt.Println(value[strit[6]] + " of " + " high")
																																																												if i == 0 {
																																																													p1power = strit[6]
																																																												} else {
																																																													p2power = strit[6]
																																																												}
																																																											}
																																																										}
																																																									}
																																																								}
																																																							}
																																																						}
																																																					}
																																																				}
																																																			}
																																																		}
																																																	}
																																																}
																																															}
																																														}
																																													}
																																												}
																																											}
																																										}
																																									}
																																								}
																																							}
																																						}
																																					}
																																				}
																																			}
																																		}
																																	}
																																}
																															}
																														}
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		flash = []string{rcsuits[2], rcsuits[3], rcsuits[4], rcsuits[5], rcsuits[6], rcsuits[7], rcsuits[8]}
		sort.Strings(flash)
		strit = []int{p2cv1, p2cv2, tc1v, tc2v, tc3v, tc4v, tc5v}
		sort.Ints(strit)
		strit2 = strit3
	}
	if p1power > p2power {
		fmt.Println("Player 1 is the winner")
	} else if p1power < p2power {
		fmt.Println("Player 2 is the winner")
	} else if p1kicker > p2kicker {
		fmt.Println("Player 1 is the winner with kicker " + value[p1kicker])
	} else if p1kicker < p2kicker {
		fmt.Println("Player 2 is the winner with kicker " + value[p2kicker])
	} else {
		fmt.Println("Draw")
	}
}
