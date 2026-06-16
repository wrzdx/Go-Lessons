package server

import (
	"CoalMine/company"
	"CoalMine/company/miner"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	company *company.Company
}

func NewHTTPHandlers(company *company.Company) *HTTPHandlers {
	return &HTTPHandlers{
		company: company,
	}
}

func (h *HTTPHandlers) HandleStartCompany(w http.ResponseWriter, r *http.Request) {
	err := h.company.Start()
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	response := StartDTO{
		Message: "Company has started successfuly",
		Started: *h.company.Started,
	}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandlers) HandleFinishCompany(w http.ResponseWriter, r *http.Request) {
	err := h.company.Finish()
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := FinishDTO{
		Message:      "Congrats",
		CompanyStats: h.company.Stats(),
	}

	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	response := struct{ Balance int }{int(h.company.Balance.Load())}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetEquipments(w http.ResponseWriter, r *http.Request) {
	response := struct{ Equipments []company.Equipment }{h.company.EquipmentInfo()}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetAllStaff(w http.ResponseWriter, r *http.Request) {
	response := struct{ Staff []miner.MinerInfo }{h.company.StaffInfo()}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetActiveStaff(w http.ResponseWriter, r *http.Request) {
	response := struct{ ActiveStaff []miner.MinerInfo }{h.company.ActiveStaffInfo()}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetMinersInfo(w http.ResponseWriter, r *http.Request) {
	response := struct{ Miners []miner.MinerInfo }{miner.GetMinersInfo()}
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
func (h *HTTPHandlers) HandleBuyEquipment(w http.ResponseWriter, r *http.Request) {
	var equipmentDTO struct {
		Name string
	}
	if err := json.NewDecoder(r.Body).Decode(&equipmentDTO); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.company.BuyEquipment(equipmentDTO.Name)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	b, err := json.MarshalIndent(struct{ Message string }{"Successful purchase"}, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandlers) HandleHireMiner(w http.ResponseWriter, r *http.Request) {
	var hireDTO struct {
		Class string
	}
	if err := json.NewDecoder(r.Body).Decode(&hireDTO); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}
	minersInfo, err := h.company.HireMiner(hireDTO.Class)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	b, err := json.MarshalIndent(minersInfo, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandlers) HandleGetStats(w http.ResponseWriter, r *http.Request) {
	response := h.company.Stats()
	b, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

func WriteError(w http.ResponseWriter, code int, err error) {
	dto := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	http.Error(w, dto.ToString(), code)
}
