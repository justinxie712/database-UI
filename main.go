package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init variables
var encryptedvolumeIops, encryptedvolumeThroughput, namespaceThroughput, regvolumeThroughput, regvolumeiopsThroughput, ioprofiledbThroughput, journalPXVolumeThroughput, journalThroughput, syncioThroughput []VolumeThroughput
var perf456 []Perf456
var resync, resyncWhileIO []Resync
var sysbench []Sysbench

// ---------- REST Endpoints ----------

func getEncryptedVolumeIops(w http.ResponseWriter, r *http.Request) {
	p := encryptedvolumeIops
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getEncryptedVolume(w http.ResponseWriter, r *http.Request) {
	p := encryptedvolumeThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getNamespaceRows(w http.ResponseWriter, r *http.Request) {
	p := namespaceThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getIoprofiledbRows(w http.ResponseWriter, r *http.Request) {
	p := ioprofiledbThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getRegVolumeIOPSRows(w http.ResponseWriter, r *http.Request) {
	p := regvolumeiopsThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getJournalPXVolumeRows(w http.ResponseWriter, r *http.Request) {
	p := journalPXVolumeThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getJournalRows(w http.ResponseWriter, r *http.Request) {
	p := journalThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getSyncioRows(w http.ResponseWriter, r *http.Request) {
	p := syncioThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getRegVolumeRows(w http.ResponseWriter, r *http.Request) {
	p := regvolumeThroughput
	t, _ := template.ParseFiles("volumeTable.html", "header.html", "volumeTableHead.html", "volumeTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "volumeTableGraphButtons.html", "volumeTableGraphs.html")
	t.Execute(w, p)
}

func getPerf456Rows(w http.ResponseWriter, r *http.Request) {
	p := perf456
	t, _ := template.ParseFiles("perf456Table.html", "perf456Header.html", "perf456TableHead.html", "perf456TableRows.html", "tableHandler.html", "dates.html", "graphHandler.html", "perf456TableGraphButtons.html", "perf456TableGraphs.html")
	t.Execute(w, p)
}

func getResyncRows(w http.ResponseWriter, r *http.Request) {
	p := resync
	t, _ := template.ParseFiles("resyncTable.html", "header.html", "resyncTableHead.html", "resyncTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "resyncTableGraphButtons.html", "resyncTableGraphs.html")
	t.Execute(w, p)
}

func getResyncWhileIORows(w http.ResponseWriter, r *http.Request) {
	p := resyncWhileIO
	t, _ := template.ParseFiles("resyncTable.html", "header.html", "resyncTableHead.html", "resyncTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "resyncTableGraphButtons.html", "resyncTableGraphs.html")
	t.Execute(w, p)
}

func getSysbenchRows(w http.ResponseWriter, r *http.Request) {
	p := sysbench
	t, _ := template.ParseFiles("sysbenchTable.html", "header.html", "sysbenchTableHead.html", "sysbenchTableRows.html", "tableHandler.html", "versions.html", "version-dropdown.html", "version-dropdown-right.html", "graphHandler.html", "sysbenchTableGraphButtons.html", "sysbenchTableGraphs.html")
	t.Execute(w, p)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Data
	encryptedvolumeIops = getVolumeThroughput("encryptedvolume_iops")
	encryptedvolumeThroughput = getVolumeThroughput("encryptedvolume_throughput")
	namespaceThroughput = getVolumeThroughput("namespace_throughput")
	ioprofiledbThroughput = getVolumeThroughput("regvolume_ioprofiledb_throughput")
	regvolumeiopsThroughput = getVolumeThroughput("regvolume_iops")
	journalPXVolumeThroughput = getVolumeThroughput("regvolume_journalpxvolume_throughput")
	journalThroughput = getVolumeThroughput("regvolume_journal_throughput")
	syncioThroughput = getVolumeThroughput("regvolume_syncio_throughput")
	regvolumeThroughput = getVolumeThroughput("regvolume_throughput")
	//perf456 = getPerf456("perf456_baremetal_avgthroughput")
	resync = getResync("resync")
	resyncWhileIO = getResync("resync_whileio")
	sysbench = getSysbench("sysbench")

	// Route Handlers / Endpoints
	r.HandleFunc("/api/tables/encryptedvolumeiops", getEncryptedVolumeIops)
	r.HandleFunc("/api/tables/encryptedvolumethroughput", getEncryptedVolume)
	r.HandleFunc("/api/tables/namespace", getNamespaceRows)
	r.HandleFunc("/api/tables/ioprofiledb", getIoprofiledbRows)
	r.HandleFunc("/api/tables/regvolumeiops", getRegVolumeIOPSRows)
	r.HandleFunc("/api/tables/journalPXVolumeThroughput", getJournalPXVolumeRows)
	r.HandleFunc("/api/tables/journalThroughput", getJournalRows)
	r.HandleFunc("/api/tables/syncioThroughput", getSyncioRows)
	r.HandleFunc("/api/tables/regvolumeThroughput", getRegVolumeRows)
	//r.HandleFunc("/api/tables/perf456", getPerf456Rows)
	r.HandleFunc("/api/tables/resync", getResyncRows)
	r.HandleFunc("/api/tables/resyncWhileIO", getResyncWhileIORows)
	r.HandleFunc("/api/tables/sysbench", getSysbenchRows)

	log.Fatal(http.ListenAndServe(":8000", r))
}
