CREATE TABLE ProductionStatuses
(
    Id               int primary key not null auto_increment,
    ProductionDate   datetime,
    EstimatedDate    datetime,
    LatestStatus     varchar(255),
    ProductionStages text
)engine=InnoDB;