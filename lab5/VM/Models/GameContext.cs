using Microsoft.EntityFrameworkCore;

namespace VM.Models;

public class GameContext : DbContext
{
    public DbSet<ItemType> ItemTypes { get; set; }
    public DbSet<Item> Items { get; set; }
    public DbSet<Location> Locations { get; set; }
    public DbSet<Messages> Messages { get; set; }
    public DbSet<Player> Players { get; set; }  

    public GameContext(DbContextOptions<GameContext> options) : base(options)
    {
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        // Можете добавить здесь дополнительную конфигурацию для моделей, если это необходимо
    }
}
