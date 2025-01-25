package telemetry

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitTracer(ctx context.Context, serviceNameKey, hostTrace string) (func(), context.Context) {

	/***************************************************************************
	Configura o exportador OTLP para enviar traces para o endpoint especificado.
	****************************************************************************/

	exporter, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(hostTrace),
		otlptracehttp.WithInsecure(),
	))

	if err != nil {
		log.Fatal(err)
	}

	/*********************************************************************
	Define atributos que identificam o serviço no sistema de rastreamento.
	**********************************************************************/

	resources := resource.NewWithAttributes(
		semconv.SchemaURL, //URL do esquema semântico para atributos padronizados (fornecido pelo OpenTelemetry).
		semconv.ServiceNameKey.String(serviceNameKey)) //Define o nome do serviço

	/**************************************************************************************************************************
	Configura o provedor de rastreamento (Tracer Provider), que é responsável por gerenciar os spans e enviá-los ao exportador.
	***************************************************************************************************************************/

	traceProvider := trace.NewTracerProvider( //Cria o Tracer Provider.
		trace.WithBatcher(exporter),   //Configura o exportador para enviar os dados em lotes (batch), otimizando o desempenho.
		trace.WithResource(resources), //Anexa os atributos definidos (nome do serviço, etc.) aos spans.
	)

	/******************************************************************************
	Define o Tracer Provider configurado como o provedor global para OpenTelemetry.
	Permite que qualquer parte do código use otel.Tracer para criar spans sem
	precisar configurar novamente o provedor.
	*******************************************************************************/

	otel.SetTracerProvider(traceProvider)

	/**************************************************************************************
		func() - Retorna uma função de limpeza para desligar o Tracer Provider corretamente.
		ctx - Retorna o contexto (ctx) atualizado para ser usado em outras partes do sistema.
	/**************************************************************************************/

	return (func() {
		err := traceProvider.Shutdown(ctx)
		if err != nil {
			return
		}
	}), ctx
}
