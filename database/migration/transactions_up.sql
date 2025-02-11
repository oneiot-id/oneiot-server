CREATE TABLE Transactions
(
    Id                   int primary key not null auto_increment,
    UserId               int,
    OrderId              int,
    PricingId            int,
    ProductionStatusesId int,
    DeliveryStatusesId   int,
    Status               varchar(255),
    CreatedAt            datetime,
    Complained           boolean
)engine=InnoDB;