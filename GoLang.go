package main
 
import "fmt"
 
type account interface {
    setNext(account)
    pay(int)
}
 
type Bank struct {
    balance     int
    accountType string
    next        account
}
 
type Paypal struct {
    balance     int
    accountType string
    next        account
}
 
type Bitcoin struct {
    balance     int
    accountType string
    next        account
}
 
// Bank Handler Functions
func (b *Bank) pay(p int) {
    if b.balance >= p {
        fmt.Printf("Paying $%d from %s account", p, b.accountType)
        return
    } else {
        fmt.Printf("Can not pay $%d from %s account, balance is: $%d , rederecting...\n", p, b.accountType, b.balance)
        b.next.pay(p)
    }
}
 
func (b *Bank) setNext(next account) {
    b.next = next
}
 
// Paypal Handler Functions
func (pp *Paypal) pay(p int) {
    if pp.balance >= p {
        fmt.Printf("Paying $%d from %s account", p, pp.accountType)
        return
    } else {
        fmt.Printf("Can not pay $%d from %s account, balance is: $%d , rederecting...\n", p, pp.accountType, pp.balance)
        pp.next.pay(p)
    }
}
 
func (po *Paypal) setNext(next account) {
    po.next = next
}
 
// Bitcoin Account Handler Functions
func (btc *Bitcoin) pay(p int) {
    if btc.balance >= p {
        fmt.Printf("Paying $%d from %s account", p, btc.accountType)
        return
    } else {
        fmt.Printf("Can not pay $%d from %s account, balance is: $%d , you should have invested earlier!", p, btc.accountType, btc.balance)
    }
}
 
func (btc *Bitcoin) setNext(next account) {
    btc.next = next
}
 
func main() {
    bank := &Bank{balance: 100, accountType: "RBC"}
    paypal := &Paypal{balance: 300, accountType: "PayPal"}
    bitcoin := &Bitcoin{balance: 9000, accountType: "Bitcoin"}
    bank.setNext(paypal)
    paypal.setNext(bitcoin)
 
    bank.pay(9001)
}
