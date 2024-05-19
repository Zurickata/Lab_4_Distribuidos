package main

import (
    pb "github.com/Zurickata/Lab_4_Distribuidos/proto"
    "google.golang.org/grpc"
    "log"
    "net"
    "context"
    "sync"
    "math/rand"
    "time"
)

type server struct {
    pb.UnimplementedDirectorServer
    mercenarios map[int32]*pb.Mercenario
    mu          sync.Mutex
    pisoActual  int32
    ready       int
    muertos     map[int32]bool
}

func (s *server) Preparacion(ctx context.Context, mercenario *pb.Mercenario) (*pb.Response, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.mercenarios[mercenario.Id] = mercenario
    return &pb.Response{Mensaje: "Preparado"}, nil
}

func (s *server) InicioPiso(ctx context.Context, piso *pb.Piso) (*pb.Response, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if piso.Numero != s.pisoActual+1 {
        return &pb.Response{Mensaje: "Piso no autorizado"}, nil
    }
    s.pisoActual = piso.Numero
    s.ready = 0
    log.Printf("Iniciando piso %d\n", s.pisoActual)
    return &pb.Response{Mensaje: "Piso iniciado"}, nil
}

func (s *server) DecisionesPiso(ctx context.Context, piso *pb.Piso) (*pb.Response, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    mercenario := s.mercenarios[piso.Mercenario.Id]
    if mercenario == nil {
        return &pb.Response{Mensaje: "Mercenario no encontrado"}, nil
    }

    // Lógica para evaluar la decisión del piso
    resultado := "Sobrevivió"
    switch piso.Numero{
    case 1:
        x, y := numsAleatorios()
        // log.Printf("x=%d, y=%d", x, y)
        switch piso.Decision{
        case 1:
            if randomInRange(0,100) < x { // Probabilidad X de morir
                resultado = "Murió"
                s.muertos[mercenario.Id] = true
            }
        case 2:
            if randomInRange(0,100) < y-x { // Probabilidad Y-X de morir
                resultado = "Murió"
                s.muertos[mercenario.Id] = true
            }
        case 3:
            if randomInRange(0,100) < 100-y { // Probabilidad 100-Y de morir
                resultado = "Murió"
                s.muertos[mercenario.Id] = true
            }
        }
    case 2:
        pasillo := randomInRange(1,2)
        if  piso.Decision != int32(pasillo) { // Si no elige el pasillo correcto muere
            resultado = "Murió"
            s.muertos[mercenario.Id] = true
        }
    case 3:
        numDir := randomInRange(1,15)
        if  piso.Decision != int32(numDir) { // Si no elige lo mismo que el Dir muere
            resultado = "Murió"
            s.muertos[mercenario.Id] = true
        }
        
    }
    log.Printf("Decisión del mercenario %s en el piso %d: %d - %s\n", mercenario.Nombre, piso.Numero, piso.Decision, resultado)
    return &pb.Response{Mensaje: resultado}, nil
}

func (s *server) ConsultarDoshBank(ctx context.Context, vacio *pb.Vacio) (*pb.DoshBank, error) {
    // Lógica para consultar el doshbank
    return &pb.DoshBank{Monto: 1000}, nil
}

func (s *server) ListoParaSiguientePiso(ctx context.Context, mercenario *pb.Mercenario) (*pb.Response, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.muertos[mercenario.Id] {
        return &pb.Response{Mensaje: "Mercenario muerto"}, nil
    }
    s.ready++
    if s.ready == len(s.mercenarios)-len(s.muertos) {
        //Pendiente Termino de la Mision
        log.Printf("Todos los mercenarios están listos para el siguiente piso\n")
    }
    return &pb.Response{Mensaje: "Preparado"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    log.Println("Servidor en ejecución en el puerto 50051...")

    grpcServer := grpc.NewServer()
    pb.RegisterDirectorServer(grpcServer, &server{
        mercenarios: make(map[int32]*pb.Mercenario),
        muertos:     make(map[int32]bool),
    })
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Función para generar un número aleatorio dentro de un rango
func randomInRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func numsAleatorios() (int, int) {
	// Inicializa la semilla para obtener números verdaderamente aleatorios
	rand.Seed(time.Now().UnixNano())

	// Genera dos números aleatorios entre 0 y 100
	numero1 := rand.Intn(101)
	numero2 := rand.Intn(101)

	// Asegúrate de que sean diferentes
	for numero1 == numero2 {
		numero2 = rand.Intn(101)
	}

	// Verifica cuál es mayor
	var mayor, menor int
	if numero1 > numero2 {
		mayor = numero1
		menor = numero2
	} else {
		mayor = numero2
		menor = numero1
	}
    
    return menor, mayor
}