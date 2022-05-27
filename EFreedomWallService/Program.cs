using CoreWCF;
using CoreWCF.Configuration;
using CoreWCF.Description;
using EFreedomWallService;
using Microsoft.AspNetCore.HttpLogging;

var builder = WebApplication.CreateBuilder(args);
builder.WebHost.ConfigureKestrel((context, options) =>
{
    options.AllowSynchronousIO = true;
});

// WSDL
builder.Services.AddServiceModelServices().AddServiceModelMetadata();

// Logging
// builder.Services.AddHttpLogging(logging => {
//     logging.LoggingFields = HttpLoggingFields.RequestPropertiesAndHeaders | HttpLoggingFields.RequestBody | HttpLoggingFields.ResponseBody;
// });

var app = builder.Build();
// app.UseHttpLogging();

app.UseServiceModel(builder =>
{
    builder.AddService<Service1>((serviceOptions) => {
        serviceOptions.BaseAddresses.Add(new Uri("http://localhost:5000/Service1"));
        //serviceOptions.BaseAddresses.Add(new Uri("https://localhost:5001/Service1"));
    })
        .AddServiceEndpoint<Service1, IService1>(new BasicHttpBinding(), "/basichttp")
        //.AddServiceEndpoint<Service1, IService1>(new WSHttpBinding(SecurityMode.Transport), "/WSHttps")
        ;   
});

var serviceMetadataBehavior = app.Services.GetRequiredService<ServiceMetadataBehavior>();
serviceMetadataBehavior.HttpGetEnabled = true;

app.Run();
