package model

type Voo struct {
    ID            int    `db:"id" json:"id"`
    Numero        string `db:"numero" json:"numero"`
    Companhia     string `db:"companhia" json:"companhia"`
    Origem        string `db:"origem" json:"origem"`
    Destino       string `db:"destino" json:"destino"`
    Status        string `db:"status" json:"status"`
}